package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

func NewBackingService(name string, validate ValidateService, check CheckService, err func(err error)) BackingService {
	return &service{
		kind:     name,
		validate: validate,
		check:    check,
		err:      err,
	}
}

type BackingService interface {
	GetBackingServices(name string) <-chan Service
}

type service struct {
	kind     string
	validate ValidateService
	check    CheckService
	err      func(err error)
}

func (s *service) GetBackingServices(name string) <-chan Service {
	return initBackingServicesFunc(s.kind, name, s.validate, s.check, s.err)
}

func initBackingServicesFunc(serviceKind, name string, validate ValidateService, check CheckService, errFunc func(err error)) <-chan Service {

	svcs := getCredentials(serviceKind)
	if len(svcs) == 0 {
		if errFunc != nil {
			errFunc(errors.New(fmt.Sprintf("backingservice %s config nil.", serviceKind)))
		}
		return nil
	}

	fmt.Printf("backingservices %v\n", svcs)

	c := make(chan Service, len(svcs))
	go func() {
		defer close(c)

		find := false
		for _, svc := range svcs {
			if svc.Name == name {
				fmt.Printf("find backingservice %s\n", name)
				c <- svc
				find = true
			}
		}

		if !find {
			fmt.Printf("find no backingservice %s\n", name)
		}
	}()

	return checkBackingServices(validateBackingServices(c, validate), check)
}

func validateBackingServices(sc <-chan Service, validate ValidateService) <-chan Service {
	nc := make(chan Service)

	go func() {
		defer close(nc)
		for {
			svc, ok := <-sc
			if ok {
				if validate(svc) {
					nc <- svc
				}
			} else {
				return
			}
		}
	}()

	return nc
}

func checkBackingServices(sc <-chan Service, checkFunc CheckService) <-chan Service {
	//we dont know sc length, must asign a length for not bocking return
	c := make(chan Service, 10)

	var wg sync.WaitGroup

	for {
		svc, ok := <-sc
		if ok {
			wg.Add(1)
			go func() {
				if checkFunc(svc) {
					c <- svc
				}
				wg.Done()
			}()
		} else {
			break
		}
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}

type CheckService func(svc Service) bool

type ValidateService func(svc Service) bool

func ValidateHPN(svc Service) bool {
	if len(svc.Credential.Host) == 0 || len(svc.Credential.Port) == 0 || len(svc.Credential.Name) == 0 {
		return false
	}
	return true
}

func ValidateHP(svc Service) bool {
	if len(svc.Credential.Host) == 0 || len(svc.Credential.Port) == 0 {
		return false
	}
	return true
}

func GenerateBackingServiceUrl(svc Service, param Params) string {
	return fmt.Sprint(svc.Credential) + fmt.Sprint(param)
}

const EnvKey = "VCAP_SERVICES"

func getCredentials(name string) ServiceList {
	s := os.Getenv(EnvKey)
	fmt.Println(s)
	if len(s) == 0 {
		return nil
	}

	m := new(map[string]ServiceList)
	if err := json.Unmarshal([]byte(s), m); err != nil {
		return nil
	}

	return (*m)[name]
}

func ErrorBackingService(err error) {
	if err != nil {
		log.Printf("config backingservice err %v", err)
	}
}
func FatalBackingService(err error) {
	if err != nil {
		log.Fatalf("config backingservice err %v", err)
	}
}
