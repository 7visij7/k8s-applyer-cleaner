package applyersCleaner

import (
	"k8s-applyer-cleaner/pkg/cleaner"
)

func Clean() {
		cleaner.CheckApplyer()
}
