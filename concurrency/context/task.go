package main

// У нас есть функция requestThatCostsMoney вызов которой стоит денег (например платное АПИ).
// нужно написать обертку RequestThatCostsMoney так, что бы при параллельных вызовах с одним и
// тем же аргументом, метод requestThatCostsMoney вызвался только один раз.

//func main() {
//	mo := MoneyOptimization{}
//	userID := "vasya"
//
//	go func() {
//		mo.RequestThatCostsMoney(userID)
//	}()
//
//	go func() {
//		mo.RequestThatCostsMoney(userID)
//	}()
//
//	// тут еще может быть множество параллельно запускаемых запросов
//}
//
//func requestThatCostsMoney(userID string) (response string) {
//	_ = userID
//	// функция, которая стоит денег
//	return "some"
//}
//
//type MoneyOptimization struct {
//	// 2 usages
//}
//
//func (o *MoneyOptimization) RequestThatCostsMoney(userID string) (response string) {
//	// тут пиши обертку над функцией requestThatCostsMoney
//	return "some"
//}
