#!/usr/bin/env sh

adb shell "zip -r runtime.zip \
/dev/__properties__ \
/system/usr/share/zoneinfo/tzdata \
/system/bin/linker \
/system/bin/linker64 \
/system/lib/libc.so \
/system/lib/libc++.so \
/system/lib/libdl.so \
/system/lib/liblog.so \
/system/lib/libm.so \
/system/lib/libstdc++.so \
/system/lib/libz.so \
/system/lib/libcrypto.so \
/system/lib/libssl.so \
/system/lib/libnetd_client.so \
/system/lib64/libc.so \
/system/lib64/libc++.so \
/system/lib64/libdl.so \
/system/lib64/liblog.so \
/system/lib64/libm.so \
/system/lib64/libstdc++.so \
/system/lib64/libz.so \
/system/lib64/libcrypto.so \
/system/lib64/libssl.so \
/system/lib64/libnetd_client.so"

adb pull runtime.zip runtime.zip
adb shell "rm runtime.zip"
