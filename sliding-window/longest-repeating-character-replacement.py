def characterReplacement_sliding_window(s: str, k: int) -> int:
    counts = [0] * 26
    left = 0
    max_count = 0  # lịch sử số chars phổ biến nhất trong window (non-decreasing)
    res = 0

    for right, ch in enumerate(s):
        idx = ord(ch) - ord('A')
        counts[idx] += 1
        # cập nhật max_count (chỉ tăng)
        if counts[idx] > max_count:
            max_count = counts[idx]

        # nếu cần thay nhiều hơn k -> thu nhỏ window
        window_len = right - left + 1
        if window_len - max_count > k:
            # giảm tần suất ký tự left đang giữ
            counts[ord(s[left]) - ord('A')] -= 1
            left += 1
            # CHÚ Ý: không giảm max_count ở đây (vẫn safe và nhanh)

        # res luôn lưu độ dài cửa sổ hợp lệ lớn nhất đã thấy
        res = max(res, right - left + 1)

    return res

# Ví dụ test
print(characterReplacement_sliding_window("BBBB", 2))      # 4
# print(characterReplacement_sliding_window("AABABBA", 1))   # 4


