package main

import (
	"fmt"
	"math"
	"optimizationMethodsTask4/common_funcs"
)

// quasiNewtonRank1 реализует Квазиньютоновский метод с поправкой ранга 1.
// startPoint - начальная точка.
// epsilon - точность (норма градиента).
// maxIter - максимальное количество итераций.
// lineSearchMaxAlpha - верхняя граница для поиска шага alpha.
// lineSearchTol - точность для метода золотого сечения.
// resetInterval - интервал для сброса H к единичной матрице (0 - не сбрасывать).
// Возвращает найденную точку минимума и количество итераций.
func quasiNewtonRank1(startPoint []float64, epsilon float64, maxIter int, lineSearchMaxAlpha, lineSearchTol float64, resetInterval int) ([]float64, int) {
	x := make([]float64, len(startPoint))
	copy(x, startPoint)
	iter := 0
	dim := len(startPoint)

	// H - аппроксимация обратной матрицы Гессе
	H := common_funcs.IdentityMatrix(dim) // Начинаем с единичной матрицы

	grad := common_funcs.GradF(x) // Начальный градиент

	// Основной цикл метода
	for iter < maxIter {
		gradNorm := common_funcs.VectorNorm(grad) // Норма градиента

		// Критерий остановки
		if gradNorm < epsilon {
			break
		}

		// Периодический сброс H к единичной матрице (для стабильности)
		if resetInterval > 0 && iter%resetInterval == 0 && iter > 0 {
			H = common_funcs.IdentityMatrix(dim)
			// fmt.Println("Сброс H на итерации", iter)
		}

		// 1. Вычисляем направление спуска: d = -H * grad
		direction := common_funcs.ScalarMult(-1.0, common_funcs.MatrixVectorMult(H, grad))

		// 2. Ищем шаг alpha с помощью одномерного поиска
		alpha := common_funcs.GoldenSectionSearch(x, direction, 0.0, lineSearchMaxAlpha, lineSearchTol)

		// 3. Обновляем точку: x_next = x + alpha * d
		xNext := common_funcs.VectorAdd(x, common_funcs.ScalarMult(alpha, direction))

		// 4. Вычисляем новый градиент
		gradNext := common_funcs.GradF(xNext)

		// 5. Вычисляем векторы delta и gamma для обновления H
		delta := common_funcs.ScalarMult(alpha, direction) // delta = x_next - x
		gamma := common_funcs.VectorSub(gradNext, grad)    // gamma = grad_next - grad

		// 6. Обновляем матрицу H по формуле Ранга 1
		Hgamma := common_funcs.MatrixVectorMult(H, gamma)
		deltaMinusHgamma := common_funcs.VectorSub(delta, Hgamma)
		denominator := common_funcs.DotProduct(deltaMinusHgamma, gamma)

		// Проверка знаменателя, чтобы избежать деления на ноль или нестабильности
		if math.Abs(denominator) > 1e-9 {
			outerProdTerm := common_funcs.OuterProduct(deltaMinusHgamma, deltaMinusHgamma)
			updateTerm := common_funcs.MatrixScalarMult(1.0/denominator, outerProdTerm)
			H = common_funcs.MatrixAdd(H, updateTerm) // H_next = H + updateTerm
		} else {
			// fmt.Println("Малый знаменатель при обновлении H на итерации", iter, ", пропуск обновления.")
			// Если знаменатель мал, обновление может быть нестабильным, пропускаем его
			// Или можно сбросить H к единичной матрице
			H = common_funcs.IdentityMatrix(dim)
		}

		// Переходим к следующей итерации
		x = xNext
		grad = gradNext
		iter++
	}

	// Сообщение, если достигнуто максимальное количество итераций
	if iter == maxIter {
		fmt.Println("Квазиньютоновский метод (Ранг 1) достиг максимального числа итераций.")
	}
	return x, iter // Возвращаем результат
}

func main() {
	startPoint := []float64{0.0, 0.0, 0.0} // Начальная точка 3D
	epsilon := 1e-5                        // Точность
	maxIter := 500                         // Макс. итераций
	lineSearchMaxAlpha := 1.0              // Макс. alpha для GSS
	lineSearchTol := 1e-6                  // Точность для GSS
	resetInterval := 5 * len(startPoint)   // Интервал сброса H (например, каждые 5*n итераций)

	// Вызываем метод
	minX, iterations := quasiNewtonRank1(startPoint, epsilon, maxIter, lineSearchMaxAlpha, lineSearchTol, resetInterval)
	minF := common_funcs.F(minX) // Значение функции в минимуме

	// Выводим результаты
	fmt.Println("\nКвазиньютоновский метод (Ранг 1):")
	fmt.Printf("Найденный минимум x: [%.6f, %.6f, %.6f]\n", minX[0], minX[1], minX[2])
	fmt.Printf("Значение функции в минимуме f(x): %.6f\n", minF)
	fmt.Printf("Количество итераций: %d\n", iterations)
}
