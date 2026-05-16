package main

// Напишите функцию bridge, которая вычитывает очередной канал из канала ins
// и перенаправляет данные из вычитанного канала в возвращаемый канал,
// пока вычитанный канал открыт и контекст не отменен.
// (!) Функция продолжает работу, пока канал ins открыт и контекст не отменен.

//func main() {
//	genVals := func() <-chan <-chan interface{} {
//		out := make(chan (<-chan interface{}))
//		go func() {
//			defer close(out)
//			for i := 0; i < 3; i++ {
//				stream := make(chan interface{}, 1)
//				stream <- i
//				close(stream)
//				out <- stream
//			}
//		}()
//		return out
//	}
//
//	var res []any
//	for v := range bridge(context.Background(), genVals()) {
//		res = append(res, v)
//	}
//
//	if !reflect.DeepEqual(res, []any{0, 1, 2}) {
//		panic("Wrong code")
//	}

//fmt.Println("You win",res)
//}
//
//func bridge(ctx context.Context, ins <-chan <-chan any) <-chan any {
//}
