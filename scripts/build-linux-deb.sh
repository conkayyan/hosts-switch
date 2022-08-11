#!/bin/sh

APP=hosts-switch
APPNAME="Hosts Switch"
APPDIR=${APP}_1.0.0

cd ../

mkdir -p build/linux/$APPDIR/usr/bin
mkdir -p build/linux/$APPDIR/usr/share/applications
mkdir -p build/linux/$APPDIR/usr/share/icons/hicolor/1024x1024/apps
mkdir -p build/linux/$APPDIR/usr/share/icons/hicolor/256x256/apps
mkdir -p build/linux/$APPDIR/DEBIAN

wails build
cp build/bin/hosts-switch build/linux/$APPDIR/usr/bin/${APP}

cp build/appicon.png build/linux/$APPDIR/usr/share/icons/hicolor/1024x1024/apps/${APP}.png
cp build/appicon.png build/linux/$APPDIR/usr/share/icons/hicolor/256x256/apps/${APP}.png

cat > build/linux/$APPDIR/usr/share/applications/${APP}.desktop << EOF
[Desktop Entry]
Version=1.0.0
Type=Application
Name=$APPNAME
Exec=$APP
Icon=$APP
Terminal=false
StartupWMClass=Lorca
EOF

cat > build/linux/$APPDIR/DEBIAN/control << EOF
Package: ${APP}
Version: 1.0.0
Section: base
Priority: optional
Architecture: amd64
Maintainer: conkay <conkay@foxmail.com>
Description: Hosts Switch
EOF

dpkg-deb --build build/linux/$APPDIR
rm -rf build/linux/$APPDIR
