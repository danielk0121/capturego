package main

import (
	"capturego/config"
	"capturego/utils"

	"github.com/getlantern/systray"
)

func main() {
	// 로거 초기화
	if err := utils.InitLogger(); err != nil {
		panic(err)
	}
	defer utils.CloseLogger()

	// 설정 파일 초기화
	if err := config.Init(); err != nil {
		utils.Error("설정 초기화 실패: %v", err)
		panic(err)
	}
	utils.Info("설정 로드 완료: 저장 경로=%s", config.Get().SaveDirectory)

	// 시스템 트레이 실행 (macOS 요구사항: 메인 스레드 점유)
	systray.Run(onTrayReady, onTrayExit)
}

// onTrayReady 트레이 아이콘이 준비된 후 호출되는 콜백
func onTrayReady() {
	utils.Info("CaptureGo 시작")

	systray.SetTitle("캡쳐고")
	systray.SetTooltip("CaptureGo")

	// 트레이 메뉴 구성
	mCapture := systray.AddMenuItem("캡처 시작", "듀얼 세이브 캡처")
	mScroll := systray.AddMenuItem("스크롤 캡처 시작", "스크롤 캡처")
	systray.AddSeparator()
	mSettings := systray.AddMenuItem("설정 열기", "브라우저로 설정 UI 열기")
	mSupport := systray.AddMenuItem("☕ 개발자 응원하기", "후원 페이지로 이동")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("앱 종료", "CaptureGo 종료")

	// 메뉴 이벤트 처리 (별도 고루틴)
	go func() {
		for {
			select {
			case <-mCapture.ClickedCh:
				utils.Info("트레이: 캡처 시작 클릭")
				// TODO: 듀얼 세이브 캡처 호출 (todo_05)
			case <-mScroll.ClickedCh:
				utils.Info("트레이: 스크롤 캡처 시작 클릭")
				// TODO: 스크롤 캡처 호출 (todo_06)
			case <-mSettings.ClickedCh:
				utils.Info("트레이: 설정 열기 클릭")
				// TODO: 브라우저로 localhost 열기 (todo_09)
			case <-mSupport.ClickedCh:
				utils.Info("트레이: 개발자 응원하기 클릭")
				// TODO: 후원 페이지 URL 오픈 (todo_10)
			case <-mQuit.ClickedCh:
				utils.Info("트레이: 앱 종료 클릭")
				systray.Quit()
			}
		}
	}()
}

// onTrayExit 트레이 종료 시 정리 작업
func onTrayExit() {
	utils.Info("CaptureGo 종료")
	// TODO: 웹서버 종료, 단축키 해제 (todo_08, todo_09)
}
