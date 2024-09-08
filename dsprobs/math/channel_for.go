package math

import (
	"fmt"
	"time"
)

func simplyPrint() {
	vdVolumeChan := make(chan string, 1)
	exit := make(chan string, 1)
	go FetchVirtualDriveVolumeAndEncrypt(vdVolumeChan, exit)
	// fmt.Printf(" I am Stuti, I am the best girl in the world")
	WaitForTasksCompletion(nil, vdVolumeChan)
	<-exit

}

func FetchVirtualDriveVolumeAndEncrypt(vdVolumeChan chan string, exit chan string) {
	for vdVolumeData := range vdVolumeChan {
		fmt.Printf(" Got: %s", vdVolumeData)
		fmt.Printf(" I am Stuti, I am the best girl in the world")
	}
	close(exit)
}

func WaitForTasksCompletion(vdEncryptOnly map[string]struct{}, vdVolumeChan chan string) {
	defer close(vdVolumeChan)
	// for
	vdVolumeChan <- "my volume data"
	time.Sleep(100)
}

// func WaitForTasksCompletion(aInCtx context.Context, aInDevice mos.AssetDeviceRegistration, aInExecCbResult *adcore.ExecuteCbResult,
// 	vdCreateTasks map[string]redfish.Task, vdEncryptTasks map[string]vdTaskMonitor, vdEncryptOnly map[string]struct{}, vdVolumeChan chan vdTaskMonitor) error {
// 	defer close(vdVolumeChan)

// func FetchVirtualDriveVolumeAndEncrypt(aInCtx context.Context, aInDevice mos.AssetDeviceRegistration,
// 	aInExecCbResult *adcore.ExecuteCbResult, vdVolumeChan chan vdTaskMonitor, vdreqParams vdReqParams, errs chan error) {
// 	logger := adlog.MustFromContext(aInCtx)

// 	defer func() {
// 		//logger.Info("closing channels in sender")
// 		close(errs)
// 	}()
// 	logger.Sugar().Infof("SUJAY : Inside FetchVirtualDriveVolumeAndEncrypt")
// 	for vdVolumeData := range vdVolumeChan {
// 		logger.Sugar().Infof("Received task Monitor %+v, vdReqParams %+v", vdVolumeData, vdreqParams)
// 		err := VirtualDriveTaskMonitorRequest(aInCtx, aInDevice, aInExecCbResult, &vdVolumeData)
// 		if err != nil {
// 			logger.Error("Error in vdTaskMonitorRequest", zap.Error(err))
// 			errs <- err
// 			return
// 		}

// 		logger.Sugar().Infof("Received vdId %+v, vdReqParams %+v", vdVolumeData, vdreqParams)
// 		err = EncryptVirtualDrives(aInCtx, aInDevice, aInExecCbResult, vdVolumeData, vdreqParams)
// 		if err != nil {
// 			logger.Error("Error in encrypt vd", zap.Error(err))
// 			errs <- err
// 			return
// 		}
// 	}
// }
