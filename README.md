# XRAYD

A wrapper over (AfterShip/email-verifier)[github.com/AfterShip/email-verifier] email verifier for ease of use

## QUICK SETUP

1. install golang on your system
2. run `go run xrayd.go <email>` in the directory of this project

For multiple emails:

3. remember to put your email(s) in the INPUT folder, then run:
`go run xrayd.go`
4. This basic version will verify emails and save the useful ones to the RESULTS/valid.txt file
5. You can optioinally adjust settings in the settings.ini file

## to-do

- [ ] add gui
- [ ] optimise code
- [ ] create library

NB: Make sure port 25 is not blocked from your device's IP