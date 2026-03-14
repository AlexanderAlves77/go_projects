package system

import "github.com/shirou/gopsutil/v3/net"

// GetNetworkStats retorna bytes enviados e recebidos
func GetNetworkStats() (uint64, uint64) {

	stats, err := net.IOCounters(false)

	if err != nil {
		return 0, 0
	}

	return stats[0].BytesSent, stats[0].BytesRecv
}

// GetNetworkSpeed calcula a velocidade de upload e download em bytes por segundo
func GetNetworkSpeed() (float64, float64) {
	stats1, err := net.IOCounters(false)
	if err != nil {
		return 0, 0
	}

	stats2, err := net.IOCounters(false)
	if err != nil {
		return 0, 0
	}

	sentSpeed := float64(stats2[0].BytesSent - stats1[0].BytesSent)
	recvSpeed := float64(stats2[0].BytesRecv - stats1[0].BytesRecv)

	return sentSpeed, recvSpeed
}

// BytesToMB converte bytes para megabytes
func BytesToMB(bytes float64) float64 {
	return bytes / 1024 / 1024
}
