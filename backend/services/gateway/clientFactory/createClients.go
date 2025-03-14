package client_factory

import (
	"fmt"
	"log/slog"
)

func CreateClients(factory ClientFactory, addr string, serviceTypes []ServiceType, logger *slog.Logger) (map[ServiceType]interface{}, error) {
	clients := make(map[ServiceType]interface{})

	for _, serviceType := range serviceTypes {
		client, err := factory.NewClient(addr, serviceType)
		if err != nil {
			logger.Error("Failed to create client", slog.String("service", serviceType.String()), slog.Any("error", err))
			return nil, fmt.Errorf("failed to create client for %s: %w", serviceType.String(), err)
		}
		clients[serviceType] = client
		logger.Info("Client created successfully", slog.String("service", serviceType.String()))
	}

	return clients, nil
}
