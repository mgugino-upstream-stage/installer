package ignition

import (
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/kubeconfig"
	"github.com/openshift/installer/pkg/asset/manifests"
	"github.com/openshift/installer/pkg/asset/tls"
)

// Stock is the stock of ignition assets that can be generated.
type Stock interface {
	// BootstrapIgnition is the asset that generates the bootstrap.ign ignition
	// config file for the bootstrap node.
	BootstrapIgnition() asset.Asset
	// MasterIgnition is the asset that generates the master.ign ignition config
	// files for the master nodes.
	MasterIgnition() asset.Asset
	// WorkerIgnition is the asset that generates the worker.ign ignition config
	// file for the worker nodes.
	WorkerIgnition() asset.Asset
}

// StockImpl is the implementation of the ignition asset stock.
type StockImpl struct {
	boostrap asset.Asset
	master   asset.Asset
	worker   asset.Asset
}

// EstablishStock establishes the stock of assets.
func (s *StockImpl) EstablishStock(
	installConfigStock installconfig.Stock,
	tlsStock tls.Stock,
	kubeconfigStock kubeconfig.Stock,
	manifestStock manifests.Stock,
) {
	s.boostrap = newBootstrap(installConfigStock, tlsStock, kubeconfigStock, manifestStock)
	s.master = newMaster(installConfigStock, tlsStock)
	s.worker = newWorker(installConfigStock, tlsStock)
}

// BootstrapIgnition returns the bootstrap asset.
func (s *StockImpl) BootstrapIgnition() asset.Asset { return s.boostrap }

// MasterIgnition returns the master asset.
func (s *StockImpl) MasterIgnition() asset.Asset { return s.master }

// WorkerIgnition returns the worker asset.
func (s *StockImpl) WorkerIgnition() asset.Asset { return s.worker }
