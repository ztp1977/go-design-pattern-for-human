package patterns

import "github.com/k0kubun/pp"

type (
	Composite struct{}

	iEmployee interface {
		getSalary() int64
	}

	employee struct {
		name   string
		salary int64
	}

	develop struct {
		*employee
	}

	designer struct {
		*employee
	}

	organization struct {
		employees []iEmployee
	}
)

func newDevelop(name string, salary int64) *develop {
	return &develop{
		&employee{
			name:   name,
			salary: salary,
		},
	}
}

func newDesigner(name string, salary int64) *designer {
	return &designer{
		&employee{
			name:   name,
			salary: salary,
		},
	}
}

func (m employee) getName() string {
	return m.name
}

func (m *employee) setSalary(salary int64) {
	m.salary = salary
}

func (m *employee) getSalary() int64 {
	return m.salary
}

func (m *organization) addEmployee(employee iEmployee) {
	m.employees = append(m.employees, employee)
}

func (m *organization) getNetSalaries() int64 {
	var netSalary int64

	for _, v := range m.employees {
		netSalary += v.getSalary()
	}

	return netSalary
}

func (Composite) Do() {

	desc.SetDesc("Composite", "公司是由员工构成的， 每个员工有其共性(领工资，有职责）", "统计工资，添加员工", "员工实例的不同", "表达部分和整体的关系", "想象一个树状结构，每个大组织都是由小的个体组成的")
	desc.print()
	o := new(organization)
	o.addEmployee(newDevelop("develop john", 100))
	o.addEmployee(newDesigner("designer john", 10000))
	pp.Println(o.getNetSalaries())
}
