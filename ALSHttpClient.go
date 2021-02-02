package animatedledstrip

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type aLSHttpClient struct {
	IpAddress string
}

func ALSHttpClient(ipAddress string) *aLSHttpClient {
	return &aLSHttpClient{IpAddress: ipAddress}
}

func (c *aLSHttpClient) resolvePath(path string) string {
	return fmt.Sprintf("http://%s:8080%s", c.IpAddress, path)
}

func (c *aLSHttpClient) get(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("GET to %s failed with %d", c.IpAddress, resp.StatusCode))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (c *aLSHttpClient) post(path string, body io.Reader) ([]byte, error) {
	resp, err := http.Post(path, "application/json", body)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("POST to %s failed with %d", c.IpAddress, resp.StatusCode))
	}
	log.Print(resp.StatusCode)
	defer resp.Body.Close()
	returnBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else {
		return returnBody, nil
	}
}

func (c *aLSHttpClient) delete(path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("POST to %s failed with %d", c.IpAddress, resp.StatusCode))
	}
	defer resp.Body.Close()
	returnBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else {
		return returnBody, nil
	}
}

func (c *aLSHttpClient) GetAnimationInfo(name string) (*animationInfo, error) {
	info, err := c.get(c.resolvePath(fmt.Sprintf("/animation/%s", name)))
	if err != nil {
		return nil, err
	}
	var newInfo animationInfo
	err = json.Unmarshal(info, &newInfo)
	if err != nil {
		return nil, err
	} else {
		return &newInfo, nil
	}
}

func (c *aLSHttpClient) GetSupportedAnimationsNames() ([]string, error) {
	names, err := c.get(c.resolvePath("/animations/names"))
	if err != nil {
		return nil, err
	}
	var newNames []string
	err = json.Unmarshal(names, &newNames)
	if err != nil {
		return nil, err
	} else {
		return newNames, nil
	}
}

func (c *aLSHttpClient) GetSupportedAnimations() ([]*animationInfo, error) {
	infos, err := c.get(c.resolvePath("/animations"))
	if err != nil {
		return nil, err
	}
	var newInfos []*animationInfo
	err = json.Unmarshal(infos, &newInfos)
	if err != nil {
		return nil, err
	} else {
		return newInfos, nil
	}
}

func (c *aLSHttpClient) GetSupportedAnimationsMap() (map[string]*animationInfo, error) {
	infos, err := c.get(c.resolvePath("/animations/map"))
	if err != nil {
		return nil, err
	}
	var newInfos map[string]*animationInfo
	err = json.Unmarshal(infos, &newInfos)
	if err != nil {
		return nil, err
	} else {
		return newInfos, nil
	}
}

func (c *aLSHttpClient) GetRunningAnimations() (map[string]*runningAnimationParams, error) {
	params, err := c.get(c.resolvePath("/running"))
	if err != nil {
		return nil, err
	}
	var newParams map[string]*runningAnimationParams
	err = json.Unmarshal(params, &newParams)
	if err != nil {
		return nil, err
	} else {
		return newParams, nil
	}
}

func (c *aLSHttpClient) GetRunningAnimationsIds() ([]string, error) {
	names, err := c.get(c.resolvePath("/running/ids"))
	if err != nil {
		return nil, err
	}
	var newNames []string
	err = json.Unmarshal(names, &newNames)
	if err != nil {
		return nil, err
	} else {
		return newNames, nil
	}
}

func (c *aLSHttpClient) GetRunningAnimationParams(id string) (*runningAnimationParams, error) {
	param, err := c.get(c.resolvePath(fmt.Sprintf("/running/%s", id)))
	if err != nil {
		return nil, err
	}
	var newParams runningAnimationParams
	err = json.Unmarshal(param, &newParams)
	if err != nil {
		return nil, err
	} else {
		return &newParams, nil
	}
}

func (c *aLSHttpClient) EndAnimation(id string) (*runningAnimationParams, error) {
	param, err := c.delete(c.resolvePath(fmt.Sprintf("/running/%s", id)))
	if err != nil {
		return nil, err
	}
	var newParams runningAnimationParams
	err = json.Unmarshal(param, &newParams)
	if err != nil {
		return nil, err
	} else {
		return &newParams, nil
	}
}

func (c *aLSHttpClient) EndAnimationFromParams(params *runningAnimationParams) (*runningAnimationParams, error) {
	return c.EndAnimation(params.Id)
}

func (c *aLSHttpClient) GetSections() ([]*section, error) {
	sects, err := c.get(c.resolvePath("/sections"))
	if err != nil {
		return nil, err
	}
	var newSects []*section
	err = json.Unmarshal(sects, &newSects)
	if err != nil {
		return nil, err
	} else {
		return newSects, nil
	}
}

func (c *aLSHttpClient) GetSectionsMap() (map[string]*section, error) {
	sects, err := c.get(c.resolvePath("/sections/map"))
	if err != nil {
		return nil, err
	}
	var newSects map[string]*section
	err = json.Unmarshal(sects, &newSects)
	if err != nil {
		return nil, err
	} else {
		return newSects, nil
	}
}

func (c *aLSHttpClient) GetSection(name string) (*section, error) {
	sect, err := c.get(c.resolvePath(fmt.Sprintf("/section/%s", name)))
	if err != nil {
		return nil, err
	}
	var newSect section
	err = json.Unmarshal(sect, &newSect)
	if err != nil {
		return nil, err
	} else {
		return &newSect, nil
	}
}

func (c *aLSHttpClient) GetFullStripSection() (*section, error) {
	return c.GetSection("fullStrip")
}

func (c *aLSHttpClient) CreateNewSection(newSection *section) (*section, error) {
	body, err := json.Marshal(newSection)
	if err != nil {
		return nil, err
	}
	sect, err := c.post(c.resolvePath("/sections"), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	var newSect section
	err = json.Unmarshal(sect, &newSect)
	if err != nil {
		return nil, err
	} else {
		return &newSect, nil
	}
}

func (c *aLSHttpClient) StartAnimation(newAnim *animationToRunParams) (*runningAnimationParams, error) {
	body, err := json.Marshal(newAnim)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	params, err := c.post(c.resolvePath("/start"), bytes.NewReader(body))
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	var newParams runningAnimationParams
	err = json.Unmarshal(params, &newParams)
	if err != nil {
		log.Print(err.Error())
		log.Print(string(params))
		return nil, err
	} else {
		return &newParams, nil
	}
}

func (c *aLSHttpClient) GetStripInfo() (*stripInfo, error) {
	info, err := c.get(c.resolvePath("/strip/info"))
	if err != nil {
		return nil, err
	}
	var newInfo stripInfo
	err = json.Unmarshal(info, &newInfo)
	if err != nil {
		return nil, err
	} else {
		return &newInfo, nil
	}
}

func (c *aLSHttpClient) GetCurrentStripColor() ([]int, error) {
	color, err := c.get(c.resolvePath("/strip/color"))
	if err != nil {
		return nil, err
	}
	var newColor []int
	err = json.Unmarshal(color, &newColor)
	if err != nil {
		return nil, err
	} else {
		return newColor, nil
	}
}

func (c *aLSHttpClient) ClearStrip() error {
	_, err := c.get(c.resolvePath("/strip/clear"))
	return err
}
