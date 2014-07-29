SHELL = /bin/bash
DESTDIR = ./dist

BUNDLE_NAME = Unicode\ Symbols
BUNDLE_VERSION = 1.0.0
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
	@install -pm 0644 ./src/Info.plist $(LBACTION)/Contents/
	@sed -e "s/\%BUNDLE_NAME\%/$(BUNDLE_NAME)/" \
	-e "s/\%BUNDLE_VERSION\%/$(BUNDLE_VERSION)/" \
	-e "s/\%BUNDLE_IDENTIFIER\%/$(BUNDLE_IDENTIFIER)/" \
	-e "s/\%BUNDLE_ICON\%/$(BUNDLE_ICON)/" \
	-e "s/\%AUTHOR\%/$(AUTHOR)/" \
	-e "s/\%TWITTER\%/$(TWITTER)/" \
	-e "s#\%WEBSITE\%#$(WEBSITE)#" \
	-e "s/\%SCRIPT_NAME\%/$(SCRIPT_NAME)/" -i "" $(LBACTION)/Contents/Info.plist
	@cp -r ./resources/* $(LBACTION)/Contents/Resources/
	go build $(LDFLAGS) -o $(LBACTION)/Contents/Scripts/$(SCRIPT_NAME) ./src

edit:
	@mvim .

refresh:
	@echo "Refreshing the LaunchBar"
	@osascript -e 'run script "tell application \"LaunchBar\" \n repeat with rule in indexing rules \n if name of rule is \"Actions\" then \n update rule \n exit repeat \n end if \n end repeat \n activate \n end tell"'

.PHONY: all dev ext clean bundle golb edit refresh
