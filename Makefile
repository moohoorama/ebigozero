.PHONY: build-aar gradle-assemble android-full-build all clean genkey

bin/game: main.go
	go build -o bin/game	

# Default target
all: android-full-build

# Build Ebiten Android AAR
build-aar:
	ebitenmobile bind -target android -androidapi 33 -javapkg com.hajimehoshi.goinovation -o mobile/android/inovation/inovation.aar ./mobile

# Gradle Assemble Android App
gradle-assemble:
	cd mobile/android && gradle :app:assembleRelease

# Android Full Build (equivalent to the compound task)
android-full-build: build-aar gradle-assemble

# Clean build artifacts
clean:
	rm -f mobile/android/inovation/inovation.aar
	cd mobile/android && gradle clean

# Generate Android keystore
genkey:
	keytool -genkeypair \
		-v \
		-keystore ./mobile/android/app/release-keystore.jks \
		-alias release \
		-keyalg RSA \
		-keysize 2048 \
		-validity 10000 \
		-storepass android123 \
		-keypass android123 \
		-dname "CN=Your Name, OU=YourOrgUnit, O=YourOrg, L=YourCity, S=YourState, C=KR"
