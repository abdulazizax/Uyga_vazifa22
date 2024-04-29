package git

import "os/exec"

func GetUserName() (string, error) {
	out, err := exec.Command("git", "config", "--global", "user.name").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func GetUserEmail() (string, error) {
    out, err := exec.Command("git", "config", "--global", "user.email").Output()
    if err != nil {
        return "", err
    }
    return string(out), nil
}
