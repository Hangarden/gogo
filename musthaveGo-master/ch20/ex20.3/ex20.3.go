// ch20/ex20.3/ex20.3.go
package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

type Sender interface { //인터페이스를 구현하면 결합을 끊을 수 있다. 결합도를 낮춤으로서 조금 더 완성도 있는 프로그램을 만들 수 있다.
	Send(parcel string) //소비자는 어찌되었든 보내거나 받기만 하면 되기 때문에 내부 구조는 알 필요가 있다. 인터페이스를 통해 구현하면 버전업에서 조금 더 자유로울 수 있다.
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	// 우체국 전송 객체를 만듭니다.
	sender := &koreaPost.PostSender{} // ❶ *koreaPost.PostSender 타입 PostSender 구조체는 send 메서드를 가지고 있기에
	SendBook("어린 왕자", sender)         // ❷ SendBook 메서드를 사용가능합니다.
	SendBook("그리스인 조르바", sender)

	fedexSender := &fedex.FedexSender{}
	SendBook("어린왕자", fedexSender)
	SendBook("어린왕자", fedexSender)
}
