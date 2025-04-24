keytool -genkeypair \
  -v \
  -keystore ./mobile/android/app/your-keystore.jks \
  -alias YOUR_KEY_ALIAS \
  -keyalg RSA \
  -keysize 2048 \
  -validity 10000 \
  -storepass YOUR_KEYSTORE_PASSWORD \
  -keypass YOUR_KEY_PASSWORD \
  -dname "CN=Your Name, OU=YourOrgUnit, O=YourOrg, L=YourCity, S=YourState, C=KR"
