SHELL = /bin/bash
DESTDIR = ./dist

BUNDLE_NAME = Unicode\ Symbols
BUNDLE_VERSION = $(shell cat VERSION)
BUNDLE_IDENTIFIER = nbjahan.launchbar.unisym
BUNDLE_ICON = at.obdev.LaunchBar:Pref_Shortcuts
AUTHOR = Nima\ Jahanshahi
TWITTER = @nbjahan
WEBSITE = https://github.com/nbjahan/launchbar-unisym
SCRIPT_NAME = unisym

LBACTION = $(DESTDIR)/$(BUNDLE_NAME).lbaction
LDFLAGS =

all: clean bundle refresh

dev: LDFLAGS := -ldflags "-X main.InDev true"
dev: all

release: all
	@echo "Making a release"
	@install -d $(DESTDIR)/release
	@ditto -ck --keepParent $(LBACTION)/ $(DESTDIR)/release/$(BUNDLE_NAME)-$(shell defaults read $(PWD)/$(LBACTION)/Contents/Info.plist CFBundleVersion).lbaction

clean:
	@$(RM) -r ./dist

bundle: golb $(LBACTION) refresh

golb:
	go install github.com/nbjahan/go-launchbar

$(LBACTION):
	@echo "Creating the $(BUNDLE_NAME).lbaction"
	@install -d ${LBACTION}/Contents/{Resources,Scripts}
	@plutil -replace CFBundleName -string $(BUNDLE_NAME) $(PWD)/src/Info.plist
	@plutil -replace CFBundleVersion -string $(BUNDLE_VERSION) $(PWD)/src/Info.plist
	@plutil -replace CFBundleIdentifier -string $(BUNDLE_IDENTIFIER) $(PWD)/src/Info.plist
	@plutil -replace CFBundleIconFile -string $(BUNDLE_ICON) $(PWD)/src/Info.plist
	@plutil -replace LBDescription.LBAuthor -string $(AUTHOR) $(PWD)/src/Info.plist
	@plutil -replace LBDescription.LBTwitter -string $(TWITTER) $(PWD)/src/Info.plist
	@plutil -replace LBDescription.LBWebsite -string $(WEBSITE) $(PWD)/src/Info.plist
	@plutil -replace LBScripts.LBDefaultScript.LBScriptName -string $(SCRIPT_NAME) $(PWD)/src/Info.plist
	@install -pm 0644 ./src/Info.plist $(LBACTION)/Contents/
	@cp -r ./resources/* $(LBACTION)/Contents/Resources/
	go build $(LDFLAGS) -o $(LBACTION)/Contents/Scripts/$(SCRIPT_NAME) ./src

edit:
	@mvim .

refresh:
	@echo "Refreshing the LaunchBar"
	@osascript -e 'run script "tell application \"LaunchBar\" \n repeat with rule in indexing rules \n if name of rule is \"Actions\" then \n update rule \n exit repeat \n end if \n end repeat \n activate \n end tell"'

.PHONY: all dev ext clean bundle golb edit refresh
