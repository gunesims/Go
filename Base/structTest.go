package Base

import "fmt"

type person struct {
	name string
	age  int
}

// 比较两个人的年龄，返回年龄大的一个，并返回年龄差
func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func personTest() {
	var tom person
	tom.name, tom.age = "tom", 12

	bob := person{"bob", 45}
	paul := person{age: 23, name: "paul"}

	tb_older, tb_diff := older(tom, bob)
	tp_older, tp_diff := older(tom, paul)
	bp_older, bp_diff := older(bob, paul)

	fmt.Printf("%s and %s, older is %s, diff is %v\n", tom.name, bob.name, tb_older.name, tb_diff)
	fmt.Printf("%s and %s, older is %s, diff is %v\n", tom.name, paul.name, tp_older.name, tp_diff)
	fmt.Printf("%s and %s, older is %s, diff is %v\n", bob.name, paul.name, bp_older.name, bp_diff)
}

type human struct {
	name   string
	age    int
	weight int
}

type student struct {
	human      human
	speciality string
}

func studentTest() {
	mark := student{human{"Mark", 16, 50}, "Computer Science"}

	fmt.Println()
	fmt.Println("His name is", mark.human.name)
	fmt.Println("His age is", mark.human.age)
	fmt.Println("His weight is", mark.human.weight)
	fmt.Println("His speciality is", mark.speciality)

}

func structTest() {
	personTest()
	studentTest()
}
