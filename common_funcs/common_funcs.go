package common_funcs

import (
	"fmt"
	"math"
)

/*// F вычисляет значение целевой функции для трехмерного вектора x.
// f(x₁, x₂, x₃) = x₁² + 2x₂² + x₁²x₂² + 2x₃ + e^(x₂² + x₃²) - x₂
func F(x []float64) float64 {
	// Проверка размерности входного вектора
	if len(x) != 3 {
		panic(fmt.Sprintf("Функция F ожидает 3-мерный вектор, получено: %d", len(x)))
	}
	x1 := x[0]
	x2 := x[1]
	x3 := x[2]
	term1 := math.Pow(x1, 2)
	term2 := 2 * math.Pow(x2, 2)
	term3 := math.Pow(x1, 2) * math.Pow(x2, 2)
	term4 := 2 * x3 // Объединенные x3 + x3
	term5 := math.Exp(math.Pow(x2, 2) + math.Pow(x3, 2))
	term6 := -x2
	return term1 + term2 + term3 + term4 + term5 + term6
}
*/
// GradF вычисляет градиент НОВОЙ целевой функции в точке x.
// Возвращает слайс []float64 размерности 3.
//func GradF(x []float64) []float64 {
//	// Проверка размерности входного вектора
//	if len(x) != 3 {
//		panic(fmt.Sprintf("Функция GradF ожидает 3-мерный вектор, получено: %d", len(x)))
//	}
//	x1 := x[0]
//	x2 := x[1]
//	x3 := x[2]
//	grad := make([]float64, 3) // Градиент теперь 3-мерный
//
//	// Частная производная по x1: ∂f/∂x₁ = 2x₁ + 2x₁x₂²
//	grad[0] = 2*x1 + 2*x1*math.Pow(x2, 2)
//
//	// Экспоненциальный член, используемый в производных по x2 и x3
//	expTerm := math.Exp(math.Pow(x2, 2) + math.Pow(x3, 2))
//
//	// Частная производная по x2: ∂f/∂x₂ = 4x₂ + 2x₁²x₂ + 2x₂e^(x₂² + x₃²) - 1
//	grad[1] = 4*x2 + 2*math.Pow(x1, 2)*x2 + 2*x2*expTerm - 1
//
//	// Частная производная по x3: ∂f/∂x₃ = 2 + 2x₃e^(x₂² + x₃²)
//	grad[2] = 2 + 2*x3*expTerm
//
//	return grad
//}

// f(x₁, x₂, x₃) = 2x₁⁴ + x₂⁴ + x₁²x₂² + x₃⁴ + x₁²x₃² + x₁ + x₂
func F(x []float64) float64 {
	// Проверка размерности входного вектора
	if len(x) != 3 {
		panic(fmt.Sprintf("Функция F (17.164) ожидает 3-мерный вектор, получено: %d", len(x)))
	}
	x1 := x[0]
	x2 := x[1]
	x3 := x[2]

	// Разбиваем формулу на отдельные слагаемые для читаемости и отладки
	term1 := 2 * math.Pow(x1, 4)               // 2x₁⁴
	term2 := math.Pow(x2, 4)                   // x₂⁴
	term3 := math.Pow(x1, 2) * math.Pow(x2, 2) // x₁²x₂²
	term4 := math.Pow(x3, 4)                   // x₃⁴
	term5 := math.Pow(x1, 2) * math.Pow(x3, 2) // x₁²x₃²
	term6 := x1                                // x₁
	term7 := x2                                // x₂

	// Возвращаем сумму всех слагаемых
	return term1 + term2 + term3 + term4 + term5 + term6 + term7
}

func GradF(x []float64) []float64 {
	if len(x) != 3 {
		panic(fmt.Sprintf("Функция GradF (17.164) ожидает 3-мерный вектор, получено: %d", len(x)))
	}
	x1 := x[0]
	x2 := x[1]
	x3 := x[2]
	grad := make([]float64, 3)
	grad[0] = 8*math.Pow(x1, 3) + 2*x1*math.Pow(x2, 2) + 2*x1*math.Pow(x3, 2) + 1
	grad[1] = 4*math.Pow(x2, 3) + 2*math.Pow(x1, 2)*x2 + 1
	grad[2] = 4*math.Pow(x3, 3) + 2*math.Pow(x1, 2)*x3
	return grad
}

// VectorNorm вычисляет евклидову норму (длину) вектора v.
func VectorNorm(v []float64) float64 {
	sumSq := 0.0
	for _, val := range v {
		sumSq += val * val
	}
	return math.Sqrt(sumSq)
}

// VectorSub выполняет вычитание векторов: res = a - b.
func VectorSub(a, b []float64) []float64 {
	if len(a) != len(b) {
		panic("Размерности векторов должны совпадать для вычитания")
	}
	res := make([]float64, len(a))
	for i := range a {
		res[i] = a[i] - b[i]
	}
	return res
}

// VectorAdd выполняет сложение векторов: res = a + b.
func VectorAdd(a, b []float64) []float64 {
	if len(a) != len(b) {
		panic("Размерности векторов должны совпадать для сложения")
	}
	res := make([]float64, len(a))
	for i := range a {
		res[i] = a[i] + b[i]
	}
	return res
}

// ScalarMult выполняет умножение вектора v на скаляр s: res = s * v.
func ScalarMult(s float64, v []float64) []float64 {
	res := make([]float64, len(v))
	for i := range v {
		res[i] = s * v[i]
	}
	return res
}

// --- Одномерный поиск методом Золотого сечения ---
// Находит alpha, минимизирующее f(x + alpha*direction) в интервале [a, b]
func GoldenSectionSearch(x, direction []float64, a, b, tol float64) float64 {
	phi := (1 + math.Sqrt(5)) / 2
	resPhi := 2 - phi
	x1 := a + resPhi*(b-a)
	x2 := b - resPhi*(b-a)
	f1 := F(VectorAdd(x, ScalarMult(x1, direction))) // Используем F из этого же пакета
	f2 := F(VectorAdd(x, ScalarMult(x2, direction))) // Используем F из этого же пакета

	for math.Abs(b-a) > tol {
		if f1 < f2 {
			b = x2
			x2 = x1
			f2 = f1
			x1 = a + resPhi*(b-a)
			f1 = F(VectorAdd(x, ScalarMult(x1, direction)))
		} else {
			a = x1
			x1 = x2
			f1 = f2
			x2 = b - resPhi*(b-a)
			f2 = F(VectorAdd(x, ScalarMult(x2, direction)))
		}
	}
	return (a + b) / 2
}

func DotProduct(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("Vector dimensions must match for dot product")
	}
	dot := 0.0
	for i := range a {
		dot += a[i] * b[i]
	}
	return dot
}

// --- Матричные операции ---
// Тип для представления матрицы (слайс слайсов)
type Matrix [][]float64

// NewMatrix создает новую матрицу n x m, заполненную нулями.
func NewMatrix(n, m int) Matrix {
	mat := make(Matrix, n)
	for i := range mat {
		mat[i] = make([]float64, m)
	}
	return mat
}

// IdentityMatrix создает единичную матрицу n x n.
func IdentityMatrix(n int) Matrix {
	mat := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		mat[i][i] = 1.0
	}
	return mat
}

//
//// Hessian вычисляет матрицу Гессе для НОВОЙ функции в точке x.
//func Hessian(x []float64) Matrix {
//	if len(x) != 3 {
//		panic(fmt.Sprintf("Функция Hessian ожидает 3-мерный вектор, получено: %d", len(x)))
//	}
//	x1 := x[0]
//	x2 := x[1]
//	x3 := x[2]
//	hess := NewMatrix(3, 3)
//	expTerm := math.Exp(math.Pow(x2, 2) + math.Pow(x3, 2))
//
//	hess[0][0] = 2 + 2*math.Pow(x2, 2)
//	hess[0][1] = 4 * x1 * x2
//	hess[0][2] = 0
//	hess[1][0] = hess[0][1] // Симметричная матрица
//	hess[1][1] = 4 + 2*math.Pow(x1, 2) + 2*expTerm + 4*math.Pow(x2, 2)*expTerm
//	hess[1][2] = 4 * x2 * x3 * expTerm
//	hess[2][0] = hess[0][2] // Симметричная матрица
//	hess[2][1] = hess[1][2] // Симметричная матрица
//	hess[2][2] = 2*expTerm + 4*math.Pow(x3, 2)*expTerm
//
//	return hess
//}

func Hessian(x []float64) Matrix {
	if len(x) != 3 {
		panic(fmt.Sprintf("Функция Hessian (17.164) ожидает 3-мерный вектор, получено: %d", len(x)))
	}
	x1 := x[0]
	x2 := x[1]
	x3 := x[2]
	hess := NewMatrix(3, 3)

	hess[0][0] = 24*math.Pow(x1, 2) + 2*math.Pow(x2, 2) + 2*math.Pow(x3, 2)
	hess[0][1] = 4 * x1 * x2
	hess[0][2] = 4 * x1 * x3
	hess[1][0] = hess[0][1] // Симметричная
	hess[1][1] = 12*math.Pow(x2, 2) + 2*math.Pow(x1, 2)
	hess[1][2] = 0
	hess[2][0] = hess[0][2] // Симметричная
	hess[2][1] = hess[1][2] // Симметричная
	hess[2][2] = 12*math.Pow(x3, 2) + 2*math.Pow(x1, 2)

	return hess
}

// Inverse3x3 вычисляет обратную матрицу для матрицы 3x3.
// Возвращает обратную матрицу и флаг bool (true, если обращение успешно).
// Использует метод алгебраических дополнений.
func Inverse3x3(m Matrix) (Matrix, bool) {
	if len(m) != 3 || len(m[0]) != 3 {
		panic("Inverse3x3 работает только для матриц 3x3")
	}
	inv := NewMatrix(3, 3)
	// Вычисляем определитель
	det := m[0][0]*(m[1][1]*m[2][2]-m[2][1]*m[1][2]) -
		m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) +
		m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])

	if math.Abs(det) < 1e-12 { // Проверка на вырожденность матрицы
		return inv, false // Невозможно обратить
	}

	invDet := 1.0 / det

	inv[0][0] = (m[1][1]*m[2][2] - m[2][1]*m[1][2]) * invDet
	inv[0][1] = (m[0][2]*m[2][1] - m[0][1]*m[2][2]) * invDet
	inv[0][2] = (m[0][1]*m[1][2] - m[0][2]*m[1][1]) * invDet
	inv[1][0] = (m[1][2]*m[2][0] - m[1][0]*m[2][2]) * invDet
	inv[1][1] = (m[0][0]*m[2][2] - m[0][2]*m[2][0]) * invDet
	inv[1][2] = (m[1][0]*m[0][2] - m[0][0]*m[1][2]) * invDet
	inv[2][0] = (m[1][0]*m[2][1] - m[2][0]*m[1][1]) * invDet
	inv[2][1] = (m[2][0]*m[0][1] - m[0][0]*m[2][1]) * invDet
	inv[2][2] = (m[0][0]*m[1][1] - m[1][0]*m[0][1]) * invDet

	return inv, true
}

// MatrixVectorMult умножает матрицу на вектор: res = m * v.
func MatrixVectorMult(m Matrix, v []float64) []float64 {
	rows := len(m)
	cols := len(m[0])
	if cols != len(v) {
		panic("Размерности матрицы и вектора не совпадают для умножения")
	}
	res := make([]float64, rows)
	for i := 0; i < rows; i++ {
		sum := 0.0
		for j := 0; j < cols; j++ {
			sum += m[i][j] * v[j]
		}
		res[i] = sum
	}
	return res
}

// OuterProduct вычисляет внешнее произведение двух векторов: res = a * b^T.
// Результатом является матрица.
func OuterProduct(a, b []float64) Matrix {
	n := len(a)
	m := len(b)
	res := NewMatrix(n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[i][j] = a[i] * b[j]
		}
	}
	return res
}

// MatrixAdd выполняет сложение матриц: res = a + b.
func MatrixAdd(a, b Matrix) Matrix {
	rowsA, colsA := len(a), len(a[0])
	rowsB, colsB := len(b), len(b[0])
	if rowsA != rowsB || colsA != colsB {
		panic("Размерности матриц должны совпадать для сложения")
	}
	res := NewMatrix(rowsA, colsA)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			res[i][j] = a[i][j] + b[i][j]
		}
	}
	return res
}

// MatrixScalarMult умножает матрицу на скаляр: res = s * m.
func MatrixScalarMult(s float64, m Matrix) Matrix {
	rows, cols := len(m), len(m[0])
	res := NewMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[i][j] = s * m[i][j]
		}
	}
	return res
}
