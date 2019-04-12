/**
 * My commodity trading network
 */
namespace org.example.mynetwork


asset Employee identified by employeeId {
  o String employeeId
  o String employeeName
  o String employeeSalary
  o String employeeDiv
  o Employeetype employeetype
  --> CompanyName owner
}

enum Employeetype {
  o PERMANANT
  o CONTRACT
}

participant CompanyName identified by companyID {
  o String companyID
  o String companyName
  o String companyLocation
}


transaction Sale {
  --> Employee employee
  --> CompanyName comapanyname
}


transaction SelectHitachiEmployees {
}

transaction SelectHumanEmployees {
}
