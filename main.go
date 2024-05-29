package main

import (
	"fmt"
)

type Inventory struct {
	Items []InventoryItem
}

type InventoryItem string

type Character struct {
	Name      string
	XP        int
	Inventory *Inventory
}

func (character *Character) wakeUp() {
	fmt.Printf("Ти прокидаєшся біля печери і не можеш згадати, що відбувалося до цього. Поруч із тобою лежить рюкзак із ліхтариком, ножем і пляшкою води. Вхід у печеру дуже темний і лякає тебе. Ти оглядаєшся навколо і розумієш, що тебе оточує ліс. Помічаєш, що вглиб лісу ведуть дивні сліди. Водночас, вузенька стежка веде в іншому напрямку - можливо вона допоможе тобі обійти ліс. \n1) Вирушити в ліс по слідах \n2)Піти стежкою в обхід\n")

	choice := getUserInput(2)

	switch choice {
	case 1:
		character.goToForest()
	case 2:
		character.walkAlongThePath()
	}
}

func (character *Character) goToForest() {
	fmt.Println("Ти обережно йдеш по слідах. Ліс стає усе густішим, сліди стає складніше помічати. Ти починаєш відчувати дивний запах, що стає все сильнішим. Дійшовши до кінця слідів ти розумієш, хто їх залишив - перед тобою лежало тіло невідомої тобі тварини. На диво, воно виглядало занадто добре для того, щоб так сильно тхнути. \n1) Оглянути тіло \n2) Вирушити далі")

	choice := getUserInput(2)

	if choice == 1 {
		character.exploreBody()
	} else if choice == 2 {
		character.goToTribe()
	}
}

func (character *Character) exploreBody() {
	fmt.Println("Обійшовши тіло з іншої сторони, ти помічаєш, що з шиї тварини стирчить стріла. Ти обережно дістаєш її. Ти не бачиш кров - замість неї з рани тварини тече дивна, густа прозора рідина. Може забрати стрілу, раптом вона знадобиться в якості зброї? \n1) Обережно покласти стрілу в рюкзак \n2) Викинути її та вирушити далі")

	choice := getUserInput(2)

	switch choice {
	case 1:
		character.takeArrow()
	case 2:
		character.goToTribe()
	}

}

func (character *Character) takeArrow() {
	fmt.Println("Ти кладеш стрілу на землю та знімаєш із себе рюкзак. Відкриваєш його. Трохи неуважно тягнешся за стрілою, торкаєшся її вістря. Воно залишає маленький поріз і забирає 10 пунктів здоровʼя. Треба знайти, чим можна обробити. Ти не відчуваєш ніякого болю, але вирішуєш поспішити далі, залишивши стрілу на своєму місці. Можливо, поруч є люди")

	character.incrementXP(-10)

	if character.XP <= 0 {
		fmt.Println("Ти помер")
		return
	}

	character.goToTribe()
}

func (character *Character) goToTribe() {
	fmt.Println("Ти озираєшся навколо щоб обрати, яким шляхом піти. Ти бачиш, що в лівий бік від тіла тварини веде ледь помітна стежка, яка радше є трохи притоптаною травою. Можливо, в кінці є люди. Ти вирушаєш туди і потрапляєш у табір. Ти обережно обходиш його - людей і справді немає. Але в тебе є відчуття, що вони були тут зовсім нещодавно. Біля одного з наметів ще тліє багаття. Ти оглядаєш його уважніше і помічаєш тіло такої самої тварини, яке бачив у лісі. Ти відчуваєш, що в тебе в голові вже паморочиться - треба швидше огледіти намети. Ти бачиш два намети - один цілий, поруч із багаттям. Інший виглядає розірваним та поламаним. \n1) Вирушити до першого намету \n2) Вирушити до другого намету")

	choice := getUserInput(2)
	switch choice {
	case 1:
		character.exploreFirstTent()
	case 2:
		character.exploreSecondTent()
	}
}

func (character *Character) exploreFirstTent() {
	fmt.Println("Ти підходиш до намету біля багаття та зазираєш усередину. Бачиш матрац і - о диво! - розкидані поруч ліки. Певно хтось дуже швидко шукав щось необхідне в аптечці, яку забрав із собою. Ти шукаєш серед ліків щось, що могло б тобі допомогти, але в ліках ти не розбираєшся, тому обираєш єдине знезаражуюче, яке знаєш, обробляєш рану і сідаєш перепочити. To be continued...")
}

func (character *Character) exploreSecondTent() {
	fmt.Println("Ти підходиш до зруйнованого намету та заглядаєш усередину. Побачене тебе приголомшує - всередині все в крові. Тебе починає нудити і ти знов непритомнієш. To be continued...")
}

func (character *Character) walkAlongThePath() {
	fmt.Println("Ти вирушаєш стежкою в обхід лісу. По дорозі насолоджуєшся приємними запахами природи. Через деякий час виходиш на гарну зелену галявину. У центрі галявини ти бачиш невеликий намет із багаттям біля входу. Ти підходиш до намету. Заглядаєш усередину. На власне здивування ти бачиш там песика. Він виглядає сумним і голодним. Ти шукаєш у рюкзаку щось, чим міг би його нагодувати, і знаходиш консерву та пляшку води. \n1) Дати песику мʼясо \n2) Дати песику води")

	choice := getUserInput(2)

	switch choice {
	case 1:
		character.feedDog()
	case 2:
		character.giveDogWater()
	}
}

func (character *Character) feedDog() {
	fmt.Println("Ти відкриваєш консерву та ставиш її біля песика. Спочатку він боїться, але згодом починає їсти та зʼїдає майже всю. \n1) Погладити песика \n2) Дати песику води")

	choice := getUserInput(2)

	switch choice {
	case 1:
		character.petDog()

	case 2:
		character.giveDogWater()
	}
}

func (character *Character) petDog() {
	fmt.Println("Ти обережно тягнеш руку до песика, щоб погладити його. На диво він навіть не гавкає і дозволяє тобі це зробити. Схоже, ти йому подобаєшся. Тобі треба йти далі, але ти не знаєш, коли повернеться той, хто залишив його тут. \n1) Відправитися в подорож із песиком \n2) Піти далі самому")

	choice := getUserInput(2)

	switch choice {
	case 1:
		character.travelWithDog()

	case 2:
		character.travelAlone()
	}
}

func (character *Character) giveDogWater() {
	fmt.Println("Ти відкриваєш пляшку з водою та наливаєш її у посудину, яку знайшов у наметі. Песик пʼє воду. Здається, тепер ти подобаєшся йому ще більше. Ти простягаєш руку, щоб погладити його")
	character.petDog()
}

func (character *Character) travelWithDog() {
	fmt.Println("Ти відправляєшся в подорож разом із песиком. Через деякий час стає зрозуміло, що у вас більше немає їжі. Тож тепер час навчитися полювати. Але разом зробити це буде легше. To be continued...")
}

func (character *Character) travelAlone() {
	fmt.Println("Ти довго вагаєшся, бо не знаєш, чи варто брати песика з собою. Ти боїшся, що в тебе не вистачить їжі та вирішуєш залишити тваринку чекати на свого господаря і відправляєшся в подорож сам. Через деяки час тебе ловлять бандити, і ти потрапляєш до них у полон. Можливо, разом із собакою ти зміг би захиститися від них. To  be continued...")

}

func (character *Character) incrementXP(XP int) {
	character.XP += XP
}

func getUserInput(optionsCount int) int {
	var inputValue int
	fmt.Scan(&inputValue)

	if inputValue > optionsCount || inputValue < 1 {
		fmt.Println("Такої опції немає. Спробуй знову")
		return getUserInput(optionsCount)
	}

	return inputValue
}

func main() {
	var name string

	fmt.Println("Привіт! Як тебе звуть?")
	fmt.Scan(&name)

	inventory := Inventory{
		Items: []InventoryItem{"canned_meat", "knife", "water_bottle"},
	}

	character := Character{
		Name:      name,
		XP:        50,
		Inventory: &inventory,
	}

	character.wakeUp()
}
