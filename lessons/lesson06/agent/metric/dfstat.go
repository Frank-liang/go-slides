package metric

import (
	"fmt"
	"github.com/glorydeath/nux"
	"log"
)

func GetDirDf(dir string) (ret *map[string]uint64) {

	mountPoints, err := nux.ListMountPoint()

	log.Print("mountPoints are: %v", mountPoints)

	if err != nil {
		log.Fatal(err)
		return
	}

	for idx := range mountPoints {
		var du *nux.DeviceUsage
		du, err = nux.BuildDeviceUsage(mountPoints[idx][0], mountPoints[idx][1], mountPoints[idx][2])
		if err != nil {
			log.Print(err)
			continue
		}

		log.Print("Getting mount points: %s", *du)

		if du.FsFile == dir {
			ret = &map[string]uint64{
				"total": du.BlocksAll,
				"used":  du.BlocksUsed,
				"free":  du.BlocksFree,
			}
			return
		}
	}
	return nil
}

func DeviceMetrics() (L []*MetricValue) {
	mountPoints, err := nux.ListMountPoint()

	if err != nil {
		log.Println(err)
		return
	}

	var diskTotal uint64 = 0
	var diskUsed uint64 = 0

	for idx := range mountPoints {
		var du *nux.DeviceUsage
		du, err = nux.BuildDeviceUsage(mountPoints[idx][0], mountPoints[idx][1], mountPoints[idx][2])
		if err != nil {
			log.Println(err)
			continue
		}

		diskTotal += du.BlocksAll
		diskUsed += du.BlocksUsed

		tags := fmt.Sprintf("mount=%s,fstype=%s", du.FsFile, du.FsVfstype)
		L = append(L, GaugeValue("df.bytes.total", du.BlocksAll, tags))
		L = append(L, GaugeValue("df.bytes.used", du.BlocksUsed, tags))
		L = append(L, GaugeValue("df.bytes.free", du.BlocksFree, tags))
		L = append(L, GaugeValue("df.bytes.used.percent", du.BlocksUsedPercent, tags))
		L = append(L, GaugeValue("df.bytes.free.percent", du.BlocksFreePercent, tags))
		L = append(L, GaugeValue("df.inodes.total", du.InodesAll, tags))
		L = append(L, GaugeValue("df.inodes.used", du.InodesUsed, tags))
		L = append(L, GaugeValue("df.inodes.free", du.InodesFree, tags))
		L = append(L, GaugeValue("df.inodes.used.percent", du.InodesUsedPercent, tags))
		L = append(L, GaugeValue("df.inodes.free.percent", du.InodesFreePercent, tags))

	}

	if len(L) > 0 && diskTotal > 0 {
		L = append(L, GaugeValue("df.statistics.total", float64(diskTotal)))
		L = append(L, GaugeValue("df.statistics.used", float64(diskUsed)))
		L = append(L, GaugeValue("df.statistics.used.percent", float64(diskUsed)*100.0/float64(diskTotal)))
	}

	return
}
