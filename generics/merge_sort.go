package generics

func MergeSort(a []int) (res []int) {
    if len(a) == 1 {
        return a
    }

    mid := len(a) / 2
    a_left, a_right := MergeSort(a[:mid]), MergeSort(a[mid:])
    var (
        i_left, i_right int
    )

    for i_left < len(a_left) || i_right < len(a_right) {
        if i_left == len(a_left) {
            res = append(res, a_right[i_right])
            i_right++
        } else if i_right == len(a_right) {
            res = append(res, a_left[i_left])
            i_left++
        } else {
            if a_left[i_left] <= a_right[i_right] {
                res = append(res, a_left[i_left])
                i_left++
            } else {
                res = append(res, a_right[i_right])
                i_right++
            }
        }
    }
    return
}
