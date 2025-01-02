package usecase

import (
	"context"
	"fmt"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type IUploadFileLogUseCase interface {
	UploadFileLogByExecutionID(ctx context.Context, projectID, pipelineID, executionID int64) error
}
type UploadFileLogUseCase struct {
	getProjectUseCase    IGetProjectUseCase
	thirdPartyToolPort   port.IThirdPartyToolPort
	storagePort          port.IStoragePort
	executionHistoryPort port.IExecutionHistoryPort
	qualityGatePort      port.IQualityGatePort
}

func (u UploadFileLogUseCase) UploadFileLogByExecutionID(ctx context.Context, projectID, pipelineID, executionID int64) error {
	project, err := u.getProjectUseCase.GetProjectById(ctx, projectID)
	if err != nil {
		log.Error(ctx, "GetProjectById", err)
		return err
	}
	var listLinkDownload []string
	execution, err := u.thirdPartyToolPort.GetExecutionDetail(ctx, project.ThirdPartyProjectID, pipelineID, executionID)
	if err != nil {
		log.Error(ctx, "GetExecutionDetail error", err)
		return err
	}
	go func() {
		project.HtmlUrl = "https://github.com/" + project.FullName
		outputFile := "result.json"

		// Create the Trivy command
		cmd := exec.Command("trivy", "repo", "--format", "json", "--output", outputFile, project.HtmlUrl)

		// Set command output to standard output for debugging (optional)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the command
		err = cmd.Run()
		if err != nil {
			log.Error(ctx, "Error running Trivy command", err)
			return
		}
		// Upload the result to the storage
		folderPathScan := fmt.Sprintf("scan/%s/%d/%d", project.Name, pipelineID, executionID)
		fullPathScan := fmt.Sprintf("%s/%s", folderPathScan, outputFile)
		data, err := ioutil.ReadFile(outputFile)
		if err != nil {
			log.Error(ctx, "ReadFile error", err)
			return
		}

		linkDownload, err := u.storagePort.UploadFile(ctx, fullPathScan, string(data))
		if err != nil {
			log.Error(ctx, "UploadFile error", err)
			return
		}

		listLinkDownload = append(listLinkDownload, linkDownload)
	}()
	for _, action := range execution.ActionExecutions {
		detailLog, err := u.thirdPartyToolPort.GetDetailLog(ctx, project.ThirdPartyProjectID, pipelineID, executionID, int64(action.Action.ID))
		if err != nil {
			log.Error(ctx, "GetDetailLog error", err)
			return err
		}
		folderPath := fmt.Sprintf("logs/%s/%d/%d/%d", project.Name, pipelineID, executionID, action.Action.ID)
		fileName := fmt.Sprintf("%s.txt", strings.ReplaceAll(action.Action.Name, " ", "_"))
		fullPath := fmt.Sprintf("%s/%s", folderPath, fileName)
		logDataList := []string{action.Status}
		logDataList = append(logDataList, detailLog.Log...)
		logData := strings.Join(logDataList, "\n")
		linkDownload, err := u.storagePort.UploadFile(ctx, fullPath, logData)
		if err != nil {
			log.Error(ctx, "UploadFile error", err)
			return err
		}
		listLinkDownload = append(listLinkDownload, linkDownload)
	}
	logs_file := strings.Join(listLinkDownload, ",")
	coverage, err := u.qualityGatePort.GetCoverage(ctx, project.SonarKey)
	if err != nil {
		log.Error(ctx, "GetCoverage error", err)
	}
	exhis := &entity.ExecutionHistoryEntity{
		ProjectID:         projectID,
		PipelineID:        pipelineID,
		ThirdPartyID:      executionID,
		LogsFile:          logs_file,
		ThirdPartyProject: project.ThirdPartyProjectID,
		Coverage:          coverage,
	}
	_, err = u.executionHistoryPort.CreateExecutionHistory(ctx, exhis)
	if err != nil {
		log.Error(ctx, "CreateExecutionHistory error", err)
		return err
	}
	log.Info(ctx, "UploadFileLogByExecutionID success")
	log.Info(ctx, "ListLinkDownload: ", listLinkDownload[0])
	return nil
}

func NewUploadFileLogUseCase(
	getProjectUseCase IGetProjectUseCase,
	thirdPartyToolPort port.IThirdPartyToolPort,
	storagePort port.IStoragePort,
	executionHistoryPort port.IExecutionHistoryPort,
	qualityGatePort port.IQualityGatePort,
) IUploadFileLogUseCase {
	return &UploadFileLogUseCase{
		getProjectUseCase:    getProjectUseCase,
		thirdPartyToolPort:   thirdPartyToolPort,
		storagePort:          storagePort,
		executionHistoryPort: executionHistoryPort,
		qualityGatePort:      qualityGatePort,
	}
}
