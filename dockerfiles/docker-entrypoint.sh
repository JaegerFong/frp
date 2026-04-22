#!/bin/sh
set -e

# 默认配置文件路径
FRPS_CONFIG=${FRPS_CONFIG:-"/etc/frp/frps.toml"}
PLUGIN_CONFIG=${PLUGIN_CONFIG:-"/etc/config/cloud.yaml"}
FEATURE_CONFIG=${FEATURE_CONFIG:-"/etc/config/feature.yaml"}

# 启动模式: both(默认), frps, plugin
MODE=${MODE:-"both"}

start_frps() {
    echo "Starting frps..."
    if [ -f "$FRPS_CONFIG" ]; then
        /usr/bin/frps -c "$FRPS_CONFIG" &
    else
        echo "Warning: frps config not found at $FRPS_CONFIG, starting with defaults"
        /usr/bin/frps &
    fi
}

start_plugin() {
    echo "Starting frps-plugin..."
    /usr/bin/frps-plugin -c "$PLUGIN_CONFIG" -f "$FEATURE_CONFIG" &
}

case "$MODE" in
    "frps")
        start_frps
        wait
        ;;
    "plugin")
        start_plugin
        wait
        ;;
    "both"|*)
        start_frps
        start_plugin
        wait
        ;;
esac
