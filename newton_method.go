package main

import (
	"fmt"
	"optimizationMethodsTask4/common_funcs"
)

// newtonMethod реализует модифицированный метод Ньютона с одномерным поиском шага.
// startPoint - начальная точка.
// epsilon - точность (норма градиента).
// maxIter - максимальное количество итераций.
// lineSearchMaxAlpha - верхняя граница для поиска шага alpha.
// lineSearchTol - точность для метода золотого сечения.
// Возвращает найденную точку минимума и количество итераций.
func newtonMethod(startPoint []float64, epsilon float64, maxIter int, lineSearchMaxAlpha, lineSearchTol float64) ([]float64, int) {
	x := make([]float64, len(startPoint))
	copy(x, startPoint)
	iter := 0
	dim := len(startPoint) // Размерность пространства

	if dim != 3 {
		panic("Метод Ньютона в данной реализации работает только для 3D")
	}

	// Основной цикл метода
	for iter < maxIter {
		grad := common_funcs.GradF(x)             // Градиент
		gradNorm := common_funcs.VectorNorm(grad) // Норма градиента

		// Критерий остановки
		if gradNorm < epsilon {
			break
		}

		hess := common_funcs.Hessian(x) // Вычисляем Гессиан
		// Пытаемся обратить Гессиан
		hessInv, invertible := common_funcs.Inverse3x3(hess)
		if !invertible {
			fmt.Println("Гессиан не обратим на итерации", iter, ", остановка.")
			// Можно попробовать перейти на шаг градиентного спуска в этом случае
			direction := common_funcs.ScalarMult(-1.0, grad)
			alpha := common_funcs.GoldenSectionSearch(x, direction, 0.0, lineSearchMaxAlpha, lineSearchTol)
			x = common_funcs.VectorAdd(x, common_funcs.ScalarMult(alpha, direction))
			iter++
			continue // Продолжить со следующей итерации
		}

		// Вычисляем направление Ньютона: p_k = -H^(-1) * grad
		direction := common_funcs.ScalarMult(-1.0, common_funcs.MatrixVectorMult(hessInv, grad))

		// Проверка направления спуска (должно быть < 0)
		// Если hessInv положительно определена, это условие выполнится
		dotProd := common_funcs.DotProduct(grad, direction)
		if dotProd >= 0 {
			fmt.Println("Направление Ньютона не является направлением спуска на итерации", iter, ", переключение на антиградиент.")
			// Если направление не является спуском (Гессиан не полож. определен),
			// можно использовать антиградиент
			direction = common_funcs.ScalarMult(-1.0, grad)
		}

		// Ищем оптимальный шаг alpha с помощью золотого сечения вдоль направления direction
		alpha := common_funcs.GoldenSectionSearch(x, direction, 0.0, lineSearchMaxAlpha, lineSearchTol)

		// Обновляем текущую точку: x = x + alpha * direction
		x = common_funcs.VectorAdd(x, common_funcs.ScalarMult(alpha, direction))

		iter++ // Увеличиваем счетчик итераций
	}

	// Сообщение, если достигнуто максимальное количество итераций
	if iter == maxIter {
		fmt.Println("Метод Ньютона достиг максимального числа итераций.")
	}
	return x, iter // Возвращаем результат
}

func main() {
	startPoint := []float64{0.0, 0.0, 0.0} // Начальная точка 3D
	epsilon := 1e-5                        // Точность
	maxIter := 100                         // Макс. итераций (Ньютон обычно сходится быстро)
	lineSearchMaxAlpha := 1.0              // Макс. alpha для GSS
	lineSearchTol := 1e-6                  // Точность для GSS

	// Вызываем метод
	minX, iterations := newtonMethod(startPoint, epsilon, maxIter, lineSearchMaxAlpha, lineSearchTol)
	minF := common_funcs.F(minX) // Значение функции в минимуме

	// Выводим результаты
	fmt.Println("\nМетод Ньютона (модифицированный):")
	fmt.Printf("Найденный минимум x: [%.6f, %.6f, %.6f]\n", minX[0], minX[1], minX[2])
	fmt.Printf("Значение функции в минимуме f(x): %.6f\n", minF)
	fmt.Printf("Количество итераций: %d\n", iterations)
}
