package main

import (
	"fmt"
	"math"
	"optimizationMethodsTask4/common_funcs"
)

// conjugateGradient реализует Метод сопряженных градиентов.
// startPoint - начальная точка.
// epsilon - точность (норма градиента).
// maxIter - максимальное количество итераций.
// lineSearchMaxAlpha - верхняя граница для поиска шага alpha.
// lineSearchTol - точность для метода золотого сечения.
// methodType - тип метода ("FR" для Флетчера-Ривза, "PR" для Полака-Рибьера).
// resetInterval - интервал для сброса направления d к антиградиенту (0 - не сбрасывать).
// Возвращает найденную точку минимума и количество итераций.
func conjugateGradient(startPoint []float64, epsilon float64, maxIter int, lineSearchMaxAlpha, lineSearchTol float64, methodType string, resetInterval int) ([]float64, int) {
	x := make([]float64, len(startPoint))
	copy(x, startPoint)
	iter := 0
	grad := common_funcs.GradF(x)                    // Начальный градиент
	direction := common_funcs.ScalarMult(-1.0, grad) // Начальное направление d0 = -grad0

	gradNormSq := common_funcs.DotProduct(grad, grad) // Квадрат нормы начального градиента

	// Основной цикл метода
	for iter < maxIter {
		gradNorm := math.Sqrt(gradNormSq) // Текущая норма градиента

		// Критерий остановки
		if gradNorm < epsilon {
			break
		}

		// Периодический сброс (рестарт) направления к антиградиенту
		if resetInterval > 0 && iter%resetInterval == 0 && iter > 0 {
			direction = common_funcs.ScalarMult(-1.0, grad)
			// fmt.Println("Сброс направления CG на итерации", iter)
		}

		// 1. Ищем шаг alpha с помощью одномерного поиска
		alpha := common_funcs.GoldenSectionSearch(x, direction, 0.0, lineSearchMaxAlpha, lineSearchTol)

		// 2. Обновляем точку: x_next = x + alpha * d
		xNext := common_funcs.VectorAdd(x, common_funcs.ScalarMult(alpha, direction))

		// 3. Вычисляем новый градиент
		gradNext := common_funcs.GradF(xNext)
		gradNormSqNext := common_funcs.DotProduct(gradNext, gradNext) // Квадрат нормы нового градиента

		// 4. Вычисляем beta по выбранной формуле
		beta := 0.0
		if gradNormSq > 1e-12 { // Избегаем деления на ноль
			switch methodType {
			case "FR": // Флетчер-Ривз
				beta = gradNormSqNext / gradNormSq
			case "PR": // Полак-Рибьер
				// beta = dot(grad_next, grad_next - grad) / dot(grad, grad)
				beta = common_funcs.DotProduct(gradNext, common_funcs.VectorSub(gradNext, grad)) / gradNormSq
				// Часто используют max(0, beta) для Полака-Рибьера для улучшения сходимости
				if beta < 0 {
					beta = 0
				}
			default:
				panic("Неизвестный тип метода сопряженных градиентов: " + methodType)
			}
		}

		// 5. Обновляем направление: d_next = -grad_next + beta * d
		direction = common_funcs.VectorAdd(common_funcs.ScalarMult(-1.0, gradNext), common_funcs.ScalarMult(beta, direction))

		// Переходим к следующей итерации
		x = xNext
		grad = gradNext
		gradNormSq = gradNormSqNext // Обновляем квадрат нормы
		iter++
	}

	// Сообщение, если достигнуто максимальное количество итераций
	if iter == maxIter {
		fmt.Printf("Метод сопряженных градиентов (%s) достиг максимального числа итераций.\n", methodType)
	}
	return x, iter // Возвращаем результат
}

func main() {
	startPoint := []float64{0.0, 0.0, 0.0} // Начальная точка 3D
	epsilon := 1e-5                        // Точность
	maxIter := 1000                        // Макс. итераций
	lineSearchMaxAlpha := 1.0              // Макс. alpha для GSS
	lineSearchTol := 1e-6                  // Точность для GSS
	resetInterval := 5 * len(startPoint)   // Интервал сброса (например, каждые 5*n итераций)

	// --- Запуск Флетчера-Ривза ---
	minX_FR, iterations_FR := conjugateGradient(startPoint, epsilon, maxIter, lineSearchMaxAlpha, lineSearchTol, "FR", resetInterval)
	minF_FR := common_funcs.F(minX_FR)
	fmt.Println("\nМетод сопряженных градиентов (Флетчер-Ривз):")
	fmt.Printf("Найденный минимум x: [%.6f, %.6f, %.6f]\n", minX_FR[0], minX_FR[1], minX_FR[2])
	fmt.Printf("Значение функции в минимуме f(x): %.6f\n", minF_FR)
	fmt.Printf("Количество итераций: %d\n", iterations_FR)

	// --- Запуск Полака-Рибьера ---
	minX_PR, iterations_PR := conjugateGradient(startPoint, epsilon, maxIter, lineSearchMaxAlpha, lineSearchTol, "PR", resetInterval)
	minF_PR := common_funcs.F(minX_PR)
	fmt.Println("\nМетод сопряженных градиентов (Полак-Рибьер):")
	fmt.Printf("Найденный минимум x: [%.6f, %.6f, %.6f]\n", minX_PR[0], minX_PR[1], minX_PR[2])
	fmt.Printf("Значение функции в минимуме f(x): %.6f\n", minF_PR)
	fmt.Printf("Количество итераций: %d\n", iterations_PR)
}
