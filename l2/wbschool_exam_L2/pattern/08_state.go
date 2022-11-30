package main

import "fmt"

/*
	Состояние.
	Вид: Поведенческий.
	Суть паттерна - позволяет объектам менять поведение в зависимости от своего состояния

	+: Избавляет от множества больших условных операторов машины состояний
	+: Концентрирует в одном месте код, связанный с определённым состоянием
	+: Упрощает код контекста
	-: Может неоправданно усложнить код, если состояний мало и они редко меняются
*/

type User struct {
	isAdmin  bool
	isAuthor bool
}

type Document struct {
	draft      IState
	moderating IState
	published  IState
	state      IState
	text       string
}

func (d *Document) render(user *User) {
	d.state.render(user)
}

func (d *Document) publish(user *User) {
	d.state.publish(user)
}

func (d *Document) setState(state IState) {
	d.state = state
}

func newDocument(text string) *Document {
	doc := &Document{text: text}
	draft := &Draft{document: doc}
	moderating := &Moderating{document: doc}
	published := &Published{document: doc}

	doc.draft = draft
	doc.moderating = moderating
	doc.published = published
	doc.setState(draft)

	return doc
}

type IState interface {
	render(user *User)
	publish(user *User)
}

type Draft struct {
	document *Document
}

func (d *Draft) render(user *User) {
	if user.isAdmin || user.isAuthor {
		fmt.Println(d.document.text)
	}
}

func (d *Draft) publish(user *User) {
	if user.isAdmin {
		d.document.setState(d.document.published)
	} else if user.isAuthor {
		d.document.setState(d.document.moderating)
	} else {
		fmt.Println("You have no required rights for that action")
	}
}

type Moderating struct {
	document *Document
}

func (m *Moderating) render(user *User) {
	if user.isAdmin {
		fmt.Println("Do you accept this text and will to publish it?")
		fmt.Println(m.document.text)
	} else if user.isAuthor {
		fmt.Println("Your text is currently under moderation")
		fmt.Println(m.document.text)
	} else {
		fmt.Println("You have no required rights for that action")
	}
}

func (m *Moderating) publish(user *User) {
	if user.isAdmin {
		m.document.setState(m.document.published)
	} else {
		fmt.Println("You have no required rights for that action")
	}
}

type Published struct {
	document *Document
}

func (p *Published) render(user *User) {
	fmt.Println(p.document.text)
}

func (p *Published) publish(user *User) {
	fmt.Println("The document is already published")
}

func main() {
	myDoc := newDocument("Hello world")

	admin := &User{isAuthor: false, isAdmin: true}
	author := &User{isAuthor: true, isAdmin: false}
	noname := &User{isAuthor: false, isAdmin: false}

	myDoc.publish(author)
	myDoc.render(admin)
	myDoc.render(noname)
	myDoc.publish(admin)
	myDoc.render(noname)
}
