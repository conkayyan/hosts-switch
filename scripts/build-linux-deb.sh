#!/bin/sh
release=$(printf %s "$1" | tr -d "v")
app=$(printf %s "$2")
app_name=$(printf %s "$3")
app_dir=${app}-linux-amd64

cd ../

mkdir -p build/linux/${app_dir}/usr/bin
mkdir -p build/linux/${app_dir}/usr/share/applications
mkdir -p build/linux/${app_dir}/usr/share/icons/hicolor/1024x1024/apps
mkdir -p build/linux/${app_dir}/usr/share/icons/hicolor/256x256/apps
mkdir -p build/linux/${app_dir}/DEBIAN

cp build/bin/${app}-linux-amd64 build/linux/${app_dir}/usr/bin/${app}

cp build/appicon.png build/linux/${app_dir}/usr/share/icons/hicolor/1024x1024/apps/${app}.png
cp build/appicon.png build/linux/${app_dir}/usr/share/icons/hicolor/256x256/apps/${app}.png

cat > build/linux/${app_dir}/usr/share/applications/${app}.desktop << EOF
[Desktop Entry]
Version=${release}
Type=Application
Name=${app_name}
Exec=${app}
Icon=${app}
Terminal=false
StartupWMClass=Lorca
EOF

cat > build/linux/${app_dir}/DEBIAN/control << EOF
Package: ${app}
Version: ${release}
Section: base
Priority: optional
Architecture: amd64
Maintainer: conkay <conkay@foxmail.com>
Description: ${app_name}
Homepage: https://github.com/conkayyan/${app}
EOF

dpkg-deb --build build/linux/${app_dir}
mv build/linux/${app_dir}.deb build/bin/
rm -rf build/linux
