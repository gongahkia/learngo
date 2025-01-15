package versionator

import (
    "runtime"
)

func CurrentVersion() string {
    return runtime.Version();
}
