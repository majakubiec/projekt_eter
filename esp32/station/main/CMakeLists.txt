idf_component_register(SRCS "station_example_main.c"
                    INCLUDE_DIRS ".")

# set WIFI_SSID and WIFI_PASSWORD env variables

if (NOT DEFINED ENV{WIFI_SSID})
    MESSAGE (FATAL_ERROR "WIFI_SSID missing; pls set this env")
endif()

if (NOT DEFINED ENV{WIFI_PASSWORD})
    MESSAGE (FATAL_ERROR "WIFI_PASSWORD missing; pls set this env")
endif()

string(REPLACE "#" "\\x23" WIFI_PASSWORD "$ENV{WIFI_PASSWORD}")

add_compile_definitions(WIFI_SSID="$ENV{WIFI_SSID}")
add_compile_definitions(WIFI_PASSWORD="${WIFI_PASSWORD}")
