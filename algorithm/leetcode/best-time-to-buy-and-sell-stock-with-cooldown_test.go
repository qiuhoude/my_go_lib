package leetcode

import "testing"

// 309. 最佳买卖股票时机含冷冻期 https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

/*
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
示例:

输入: [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

思路:
自顶向下的思路思考, 第一天分 买入 或 不买入两种情况,分别递归的往下进行发展
定义函数买入 f(input) 买入的最大收益
定义函数 g(remainIndex) 在余下的天数收益最大值
最后的结果是: max(f(0),g(1))

思路2:
动态规划的思路,将分成3种状态
f[i][0] 第i天，当前已经持有一只股票的最大收益 i-1天或之前就已经持有
f[i][1] 第i天，当前没有持有一只股票的最大收益(冷冻期) i+1天才能买入
f[i][2] 第i天，当前没有持有一只股票的最大收益(非冷冻期)
状态转换=>
f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])	// 要变成持有一只股票并且收益最大，需要比较 前一天已经股票状态，前一天刚刚买过一只股票状态
f[i][1] = f[i-1][0] + prices[i]					// 要变成冷冻期，前一天卖掉了收益的股票才行
f[i][2] = max(f[i−1][1],f[i−1][2])

初始化
f[0][0] = -prices[0] // 第一天就持有股票所收益是负数
f[0][0] = 0
f[0][0] = 0
*/

// 动态规划的方式
func maxProfit2(prices []int) int {
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = maxFn(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i] // 前一天卖掉了收益的股票,i天处于冷冻期
		dp[i][2] = maxFn(dp[i-1][1], dp[i-1][2])
	}

	return maxFn(dp[n-1][2], maxFn(dp[n-1][1], dp[n-1][0])) // 最后一天是买入状态必然是亏的,所以可以不比较
}

// 递归回溯+记忆化搜索
func maxProfit(prices []int) int {
	n := len(prices)
	var fFn func(int) int
	var gFn func(int) int

	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const MEMORY_INI_VAL = -1
	// 记忆化搜索
	fMemory, gMemory := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fMemory[i], gMemory[i] = MEMORY_INI_VAL, MEMORY_INI_VAL
	}

	// in 买入下标
	fFn = func(in int) int {
		mVal := 0
		for out := in + 1; out < n; out++ { // 选择卖出的日子,选出利润最大的
			if prices[out] > prices[in] { // 卖出价大于买入价才有意义
				cur := prices[out] - prices[in] // 当前利润
				remain := 0                     // 余下天数的利润
				if out < n-3 {                  // 至少还有3天才能有一个交易周期
					if gMemory[out+2] != MEMORY_INI_VAL {
						remain = gMemory[out+2]
					} else {
						remain = gFn(out + 2)
					}
				}
				// 当前利润+余下天数的利润
				mVal = maxFn(cur+remain, mVal)
			}
		}
		fMemory[in] = mVal
		return mVal
	}

	gFn = func(remainIndex int) int {
		if remainIndex > n-2 {
			return 0
		}
		// remainIndex 买入,或不买入
		fVal := 0
		if fMemory[remainIndex] != MEMORY_INI_VAL {
			fVal = fMemory[remainIndex]
		} else {
			fVal = fFn(remainIndex)
		}
		gVal := 0
		if gMemory[remainIndex+1] != MEMORY_INI_VAL {
			gVal = gMemory[remainIndex+1]
		} else {
			gVal = gFn(remainIndex + 1)
		}
		mVal := maxFn(fVal, gVal)
		gMemory[remainIndex] = mVal
		return mVal
	}
	res := gFn(0)
	return res
}

func Test_maxProfit(t *testing.T) {

	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{"1", []int{1, 2, 3, 0, 2}, 3},
		{"2", []int{3, 7, 6, 0, 5, 4}, 9},
		{"3", []int{2, 4, 1, 7, 11}, 10},
		{"4", []int{106, 373, 495, 46, 359, 919, 906, 440, 783, 583, 784, 73, 238, 701, 972, 308, 165, 774, 990, 675, 737, 990, 713, 157, 211, 880, 961, 132, 980, 136, 285, 239, 628, 221, 948, 939, 28, 541, 414, 180, 171, 640, 297, 873, 59, 814, 832, 611, 868, 633, 101, 67, 396, 264, 445, 548, 257, 656, 624, 71, 607, 67, 836, 14, 373, 205, 434, 203, 661, 793, 45, 623, 140, 67, 177, 885, 155, 764, 363, 269, 599, 32, 228, 111, 102, 565, 918, 592, 604, 244, 982, 533, 781, 604, 115, 429, 33, 894, 778, 885, 145, 888, 577, 275, 644, 824, 277, 302, 182, 94, 479, 563, 52, 771, 544, 794, 964, 827, 744, 366, 548, 761, 477, 434, 999, 86, 1000, 5, 99, 311, 346, 609, 778, 937, 372, 793, 754, 191, 592, 860, 748, 297, 610, 386, 146, 220, 7, 113, 657, 438, 482, 700, 158, 884, 877, 964, 777, 139, 809, 489, 383, 92, 581, 970, 899, 947, 864, 443, 490, 825, 674, 906, 402, 270, 416, 611, 949, 476, 775, 899, 837, 796, 227, 232, 226, 11, 266, 889, 215, 6, 182, 430, 5, 706, 994, 128, 359, 841, 439, 263, 491, 689, 638, 485, 763, 695, 135, 800, 763, 54, 569, 387, 112, 316, 193, 675, 546, 531, 954, 571, 208, 282, 557, 892, 469, 875, 765, 592, 374, 276, 892, 843, 625, 180, 249, 292, 477, 882, 837, 112, 46, 667, 187, 93, 418, 790, 903, 12, 978, 510, 647, 446, 597, 958, 678, 897, 420, 907, 256, 170, 669, 920, 711, 635, 995, 259, 994, 634, 583, 175, 380, 435, 942, 739, 921, 132, 455, 986, 567, 464, 301, 10, 579, 84, 745, 717, 588, 414, 375, 319, 770, 310, 510, 521, 88, 445, 59, 460, 120, 765, 480, 441, 169, 374, 180, 947, 179, 346, 490, 417, 149, 140, 577, 624, 427, 238, 341, 686, 623, 228, 672, 859, 372, 938, 567, 141, 133, 671, 255, 997, 272, 591, 115, 340, 692, 531, 235, 123, 677, 980, 31, 774, 135, 194, 956, 723, 779, 375, 546, 59, 695, 616, 416, 362, 38, 145, 782, 184, 418, 806, 444, 177, 360, 485, 941, 998, 85, 840, 740, 545, 49, 570, 17, 824, 845, 749, 177, 727, 238, 656, 787, 425, 473, 323, 683, 578, 442, 436, 444, 595, 367, 44, 467, 93, 507, 949, 598, 579, 471, 1, 347, 982, 232, 878, 217, 845, 777, 284, 527, 529, 100, 482, 456, 814, 457, 251, 494, 419, 922, 139, 706, 384, 954, 365, 680, 70, 810, 764, 820, 992, 622, 29, 697, 294, 553, 655, 63, 934, 827, 157, 680, 812, 729, 486, 403, 151, 988, 926, 460, 193, 294, 423, 774, 715, 906, 957, 598, 929, 339, 119, 686, 88, 228, 803, 806, 743, 430, 315, 224, 712, 724, 69, 606, 411, 271, 700, 520, 179, 916, 490, 652, 319, 69, 245, 827, 185, 200, 911, 363, 335, 50, 353, 551, 737, 15, 429, 966, 766, 307, 829, 379, 184, 779, 239, 254, 904, 262, 719, 321, 380, 253, 564, 348, 878, 570, 470, 313, 752, 563, 164, 301, 239, 856, 491, 154, 795, 640, 199, 940, 420, 201, 254, 400, 865, 886, 819, 424, 292, 257, 572, 112, 590, 984, 421, 639, 705, 707, 779, 660, 4, 817, 265, 465, 737, 56, 564, 797, 178, 552, 988, 621, 98, 665, 379, 607, 300, 439, 269, 196, 94, 860, 540, 830, 756, 294, 806, 321, 930, 623, 206, 440, 730, 829, 566, 420, 488, 49, 438, 447, 294, 548, 804, 514, 45, 383, 431, 373, 424, 11, 377, 868, 559, 316, 831, 464, 211, 710, 803, 680, 665, 39, 523, 951, 219, 293, 909, 838, 708, 663, 627, 220, 100, 565, 269, 982, 236, 185, 194, 697, 556, 767, 541, 360, 103, 497, 271, 919, 19, 206, 73, 393, 50, 421, 466, 970, 329, 105, 618, 17, 687, 578, 260, 759, 366, 334, 686, 613, 616, 893, 351, 847, 861, 452, 454, 454, 88, 135, 357, 194, 220, 504, 36, 916, 246, 718, 172, 395, 292, 613, 533, 662, 983, 701, 877, 842, 445, 263, 529, 679, 526, 31, 385, 918, 898, 584, 846, 474, 648, 67, 331, 890, 174, 766, 274, 476, 414, 701, 835, 537, 531, 578, 7, 479, 906, 93, 667, 735, 435, 899, 49, 953, 854, 843, 326, 322, 13, 865, 791, 828, 686, 760, 957, 655, 601, 406, 185, 738, 788, 519, 874, 630, 440, 839, 511, 149, 715, 566, 988, 0, 354, 498, 81, 193, 335, 196, 157, 515, 590, 768, 366, 287, 386, 502, 143, 547, 659, 616, 822, 479, 813, 497, 222, 285, 6, 453, 363, 906, 388, 733, 804, 624, 963, 634, 319, 817, 674, 754, 378, 999, 373, 793, 419, 246, 274, 960, 1, 130, 186, 576, 382, 204, 227, 607, 435, 299, 790, 603, 196, 236, 955, 654, 812, 214, 297, 926, 721, 977, 568, 339, 913, 297, 621, 783, 242, 257, 483, 325, 998, 164, 586, 782, 597, 210, 522, 369, 676, 339, 626, 650, 634, 477, 793, 85, 12, 695, 655, 53, 287, 730, 0, 689, 225, 805, 593, 430, 610, 963, 172, 148, 740, 579, 16, 523, 570, 802, 627, 220, 664, 945, 788, 500, 90, 410, 916, 481, 454, 538, 622, 161, 373, 523, 757, 446, 855, 958, 390, 333, 927, 253, 814, 442, 77, 325, 14, 655, 502, 200, 791, 58, 714, 951, 370, 557, 261, 859, 199, 46, 775, 249, 369, 233, 321, 733, 310, 503, 539, 618, 839, 272, 315, 999, 229, 390, 359, 528, 334, 878, 342, 977, 869, 704, 564, 506, 867, 77, 248, 674, 557, 258, 710, 126, 617, 531, 969, 289, 578, 947, 103, 581, 599, 918, 686, 143, 253, 56, 393, 58, 144, 211, 806, 285, 635, 203, 194, 884, 687, 653, 856, 688, 623, 568, 394, 749, 302, 534, 631, 894, 167, 111, 227, 296, 41, 854, 81, 147, 656, 319, 748, 530, 457, 340, 223, 896, 77, 166, 974, 659, 36, 338, 177, 496, 483, 690, 569, 504, 211, 554, 758, 732, 660, 61, 62, 669, 273, 0, 616, 899, 789, 380, 386, 357, 403, 251, 926, 636, 419, 148, 820, 774, 485, 497, 370, 907, 973, 255, 277, 341, 466, 254, 333, 219, 819, 521, 974, 213, 590, 981, 697, 927, 904, 717, 726, 574, 94, 625, 991, 378, 249, 388, 786, 355, 69, 318, 357, 467, 695, 825, 585, 940, 323, 993, 549, 485, 564, 833, 530, 398, 789, 608, 59, 541, 915, 81, 681, 544, 460, 318, 954, 764, 879, 708, 258, 276, 259, 505, 649, 529, 824, 914, 660, 490, 666, 676, 618, 339, 712, 981, 802, 239, 605, 270, 29, 491, 41, 243, 361, 644, 327, 472, 460, 725, 864, 129, 142, 610, 782, 935, 929, 63, 865, 287, 316, 740, 212, 152, 567, 620, 591, 394, 805, 586, 177, 918, 516, 911, 944, 427, 128, 778, 930, 965, 27, 633, 534, 567, 575, 247, 691, 571, 775, 456, 622, 219, 698, 772, 305, 27, 810, 690, 555, 222, 877, 985, 493, 202, 84, 180, 133, 129, 539, 151, 275, 234, 999, 676, 629, 715, 839, 6, 789, 663, 467, 435, 275, 580, 296, 8, 73, 849, 456, 681, 794, 954, 543, 602, 615, 4, 131, 593, 778, 175, 587, 670, 88, 648, 79, 703, 99, 457, 261, 722, 357, 966, 724, 523, 612, 610, 376, 575, 174, 2, 53, 637, 478, 850, 250, 238, 344, 381, 543, 686, 761, 582, 598, 804, 12, 128, 928, 133, 998, 188, 598, 590, 507, 898, 402, 771, 703, 912, 744, 317, 300, 852, 631, 767, 157, 278, 520, 452, 721, 560, 112, 206, 69, 317, 498, 942, 942, 963, 347, 61, 186, 390, 128, 946, 462, 230, 551, 956, 195, 960, 143, 225, 654, 255, 370, 778, 770, 487, 192, 479, 180, 505, 509, 508, 717, 976, 826, 346, 521, 472, 148, 965, 965, 971, 421, 402, 233, 76, 543, 533, 815, 281, 986, 638, 936, 139, 754, 728, 779, 551, 425, 17, 546, 516, 862, 963, 648, 127, 510, 453, 311, 759, 654, 550, 755, 654, 567, 129, 34, 927, 900, 421, 961, 923, 117, 766, 71, 132, 680, 917, 460, 609, 874, 179, 336, 496, 287, 61, 846, 228, 871, 590, 858, 404, 646, 449, 770, 724, 245, 634, 900, 496, 157, 864, 407, 632, 998, 596, 451, 482, 921, 102, 624, 148, 346, 282, 624, 150, 523, 598, 492, 267, 54, 889, 872, 979, 38, 1, 282, 513, 877, 798, 994, 400, 254, 435, 487, 707, 459, 575, 275, 297, 165, 104, 468, 80, 820, 571, 215, 869, 381, 107, 209, 762, 455, 415, 810, 137, 674, 304, 692, 639, 304, 534, 348, 938, 575, 432, 471, 74, 631, 291, 405, 622, 352, 58, 549, 832, 655, 458, 688, 468, 827, 447, 946, 181, 908, 585, 53, 905, 733, 363, 210, 536, 960, 577, 815, 462, 193, 31, 731, 8, 538, 695, 936, 795, 139, 782, 357, 52, 492, 610, 512, 544, 323, 276, 649, 940, 54, 749, 723, 544, 365, 500, 441, 284, 17, 660, 748, 871, 701, 591, 356, 64, 34, 422, 713, 978, 96, 218, 756, 833, 177, 832, 61, 91, 764, 510, 188, 415, 622, 473, 549, 944, 716, 998, 528, 61, 829, 953, 280, 284, 706, 323, 981, 405, 91, 887, 568, 874, 725, 236, 933, 41, 895, 940, 375, 468, 314, 667, 694, 609, 631, 621, 655, 640, 835, 513, 461, 854, 419, 455, 860, 912, 572, 769, 963, 213, 818, 158, 840, 699, 414, 969, 430, 59, 855, 997, 997, 884, 349, 723, 837, 488, 430, 671, 743, 943, 310, 399, 884, 423, 486, 587, 491, 106, 716, 0, 768, 704, 483, 663, 827, 587, 915, 904, 742, 976, 6, 455, 221, 849, 920, 548, 156, 35, 101, 270, 684, 123, 549, 649, 977, 711, 965, 492, 525, 130, 744, 697, 910, 699, 301, 285, 696, 313, 117, 122, 777, 163, 789, 924, 543, 446, 60, 214, 102, 97, 45, 670, 960, 23, 522, 680, 178, 757, 792, 633, 244, 327, 129, 188, 357, 733, 419, 496, 774, 408, 90, 615, 663, 321, 526, 946, 990, 273, 135, 373, 719, 870, 810, 798, 826, 64, 971, 156, 233, 587, 253, 712, 384, 964, 173, 511, 116, 291, 639, 450, 947, 623, 656, 548, 605, 498, 709, 143, 895, 739, 663, 160, 442, 820, 802, 380, 413, 356, 742, 744, 764, 421, 355, 499, 614, 678, 336, 850, 1000, 463, 794, 388, 478, 188, 576, 822, 164, 209, 465, 901, 116, 729, 891, 952, 611, 15, 798, 731, 711, 6, 459, 587, 278, 996, 220, 642, 563, 363, 271, 16, 379, 959, 332, 315, 414, 659, 602, 786, 571, 78, 450, 544, 393, 404, 953, 480, 215, 771, 419, 8, 738, 36, 191, 138, 204, 146, 923, 413, 908, 998, 46, 928, 678, 425, 584, 372, 689, 245, 721, 177, 833, 44, 784, 121, 164, 16, 714, 680, 974, 685, 340, 810, 101, 301, 791, 716, 697, 768, 33, 901, 994, 417, 353, 248, 559, 807, 64, 450, 724, 896, 889, 880, 818, 89, 495, 848, 915, 450, 409, 958, 413, 149, 743, 782, 64, 687, 196, 737, 769, 311, 429, 598, 585, 690, 919, 331, 94, 211, 633, 888, 856, 844, 870, 931, 934, 66, 407, 121, 902, 417, 522, 423, 821, 196, 625, 855, 830, 673, 463, 181, 857, 775, 374, 490, 971, 751, 835, 823, 770, 79, 916, 80, 829, 810, 856, 674, 524, 352, 251, 548, 899, 363, 465, 0, 989, 322, 51, 86, 740, 542, 920, 310, 365, 677, 287, 688, 373, 225, 774, 331, 430, 482, 630, 46, 567, 236, 370, 502, 347, 191, 137, 646, 218, 634, 399, 278, 423, 540, 26, 612, 700, 43, 508, 176, 268, 525, 267, 676, 257, 651, 88, 349, 556, 6, 463, 29, 410, 753, 224, 693, 535, 747, 40, 854, 155, 376, 192, 434, 12, 342, 98, 718, 639, 951, 205, 923, 354, 564, 988, 960, 676, 965, 29, 104, 898, 535, 915, 868, 768, 269, 294, 944, 523, 145, 895, 382, 53, 935, 671, 518, 338, 623, 524, 204, 146, 900, 161, 258, 739, 417, 119, 825, 336, 182, 123, 749, 355, 188, 109, 740, 945, 826, 921, 123, 65, 69, 682, 461, 259, 661, 247, 523, 796, 153, 142, 851, 411, 536, 190, 478, 417, 296, 113, 158, 263, 754, 532, 368, 748, 42, 890, 129, 643, 717, 564, 525, 5, 348, 204, 383, 427, 696, 861, 684, 902, 591, 609, 467, 837, 104, 565, 168, 828, 916, 645, 232, 153, 794, 384, 642, 753, 550, 176, 142, 132, 141, 192, 635, 14, 634, 329, 403, 790, 460, 29, 512, 443, 15, 74, 114, 456, 487, 303, 13, 822, 429, 136, 113, 637, 283, 542, 519, 411, 564, 220, 346, 907, 389, 780, 479, 480, 179, 385, 285, 445, 393, 508, 885, 697, 168, 542, 357, 553, 149, 710, 126, 508, 271, 845, 689, 231, 217, 984, 848, 905, 87, 168, 1000, 169, 336, 672, 595, 501, 411, 81, 707, 708, 634, 150, 722, 379, 77, 762, 737, 585, 419, 428, 37, 869, 509, 222, 335, 192, 980, 209, 883, 864, 215, 497, 992, 155, 408, 652, 927, 990, 708, 439, 857, 934, 838, 69, 140, 713, 573, 939, 338, 628, 685, 412, 147, 530, 643, 471, 545, 58, 111, 132, 665, 572, 38, 176, 460, 555, 997, 61, 602, 471, 901, 620, 830, 577, 436, 495, 685, 619, 600, 549, 270, 77, 512, 249, 697, 466, 864, 336, 981, 901, 573, 702, 694, 937, 299, 565, 436, 613, 187, 377, 364, 473, 405, 384, 280, 658, 561, 85, 987, 302, 856, 107, 191, 486, 464, 165, 514, 948, 227, 310, 133, 799, 363, 481, 289, 153, 990, 445, 246, 454, 729, 887, 980, 546, 730, 528, 817, 521, 437, 376, 238, 965, 511, 995, 432, 227, 883, 550, 904, 818, 556, 295, 413, 786, 861, 248, 113, 660, 982, 445, 292, 562, 722, 433, 621, 783, 375, 53, 236, 856, 275, 898, 532, 915, 804, 362, 545, 373, 397, 740, 453, 726, 983, 665, 715, 379, 176, 408, 3, 911, 573, 883, 195, 254, 469, 758, 844, 355, 409, 562, 307, 752, 274, 105, 227, 635, 121, 335, 338, 46, 993, 243, 567, 765, 589, 806, 405, 558, 25, 246, 526, 490, 306, 295, 112, 847, 792, 759, 881, 500, 398, 791, 266, 33, 372, 546, 217, 286, 898, 596, 955, 720, 70, 9, 458, 698, 367, 936, 134, 95, 887, 300, 975, 72, 235, 77, 870, 943, 511, 883, 923, 619, 812, 904, 990, 643, 871, 346, 588, 807, 957, 681, 581, 195, 82, 448, 146, 807, 559, 21, 412, 950, 536, 681, 541, 856, 631, 378, 258, 736, 116, 580, 20, 606, 748, 537, 343, 681, 22, 711, 628, 536, 395, 422, 874, 135, 519, 294, 876, 185, 583, 392, 253, 220, 80, 341, 203, 970, 825, 762, 558, 942, 797, 651, 290, 8, 414, 375, 913, 167, 977, 94, 706, 970, 286, 278, 349, 909, 422, 887, 921, 492, 467, 550, 538, 555, 841, 446, 199, 312, 816, 562, 296, 609, 39, 393, 240, 763, 222, 828, 802, 944, 714, 325, 334, 936, 995, 950, 487, 433, 195, 370, 498, 926, 109, 543, 885, 463, 687, 171, 703, 985, 292, 123, 314, 174, 183, 588, 487, 857, 63, 736, 126, 156, 172, 367, 313, 672, 494, 56, 202, 470, 821, 735, 72, 812, 282, 570, 756, 633, 82, 52, 920, 300, 199, 927, 534, 214, 354, 764, 84, 419, 462, 5, 246, 787, 305, 788, 852, 58, 698, 241, 184, 904, 533, 333, 857, 215, 531, 81, 862, 567, 56, 773, 741, 169, 982, 965, 302, 724, 145, 342, 731, 184, 914, 977, 933, 727, 918, 420, 438, 491, 300, 104, 107, 730, 506, 214, 214, 968, 351, 66, 844, 965, 758, 845, 503, 495, 503, 208, 281, 622, 905, 49, 751, 660, 268, 420, 360, 354, 971, 441, 565, 513, 711, 283, 695, 109, 432, 127, 399, 177, 640, 67, 77, 364, 327, 943, 1000, 979, 278, 526, 222, 929, 120, 753, 580, 743, 456, 241, 148, 339, 599, 919, 11, 473, 101, 365, 789, 465, 819, 778, 134, 278, 89, 598, 801, 904, 681, 695, 599, 43, 897, 763, 193, 257, 719, 410, 610, 58, 72, 912, 598, 793, 347, 640, 725, 855, 390, 754, 785, 70, 449, 24, 962, 843, 735, 729, 893, 797, 512, 390, 57, 474, 336, 855, 970, 389, 722, 735, 464, 28, 894, 664, 645, 96, 357, 255, 117, 795, 479, 151, 790, 432, 748, 780, 940, 255, 204, 607, 999, 989, 48, 139, 945, 783, 736, 826, 640, 597, 171, 423, 457, 972, 424, 419, 140, 706, 333, 648, 557, 155, 707, 547, 337, 66, 338, 818, 829, 972, 13, 500, 310, 961, 668, 991, 407, 386, 893, 589, 308, 129, 739, 689, 452, 361, 822, 418, 606, 961, 981, 385, 132, 7, 938, 102, 942, 534, 154, 133, 921, 51, 257, 205, 281, 771, 215, 508, 816, 466, 744, 85, 141, 163, 418, 894, 386, 779, 142, 137, 825, 556, 764, 647, 414, 792, 605, 945, 36, 427, 173, 907, 356, 893, 875, 449, 621, 181, 963, 801, 14, 502, 234, 495, 437, 86, 635, 846, 182, 182, 540, 340, 648, 772, 195, 93, 539, 716, 573, 431, 342, 989, 156, 745, 436, 709, 22, 532, 100, 504, 0, 985, 838, 461, 725, 555, 219, 710, 568, 914, 736, 791, 507, 615, 442, 494, 977, 546, 519, 389, 614, 78, 172, 991, 255, 154, 243, 495, 876, 267, 948, 657, 692, 46, 107, 864, 168, 785, 965, 740, 16, 878, 713, 79, 517, 68, 208, 621, 13, 362, 99, 379, 109, 823, 960, 645, 440, 944, 342, 710, 267, 656, 646, 639, 453, 155, 867, 456, 606, 328, 444, 136, 89, 104, 650, 36, 240, 320, 31, 352, 522, 520, 260, 510, 981, 591, 655, 668, 23, 544, 320, 541, 707, 133, 708, 809, 972, 196, 59, 383, 642, 153, 993, 837, 98, 300, 751, 564, 399, 848, 325, 903, 534, 662, 201, 690, 300, 404, 115, 104, 600, 236, 752, 651, 640, 244, 254, 40, 549, 304, 86, 600, 755, 59, 662, 106, 290, 368, 725, 138, 705, 28, 550, 955, 277, 959, 346, 721, 759, 569, 420, 424, 59, 989, 438, 867, 725, 544, 178, 575, 137, 21, 536, 72, 617, 194, 421, 226, 378, 483, 880, 688, 791, 930, 97, 831, 113, 711, 445, 308, 813, 967, 120, 769, 329, 718, 899, 364, 638, 308, 644, 25, 138, 88, 732, 922, 721, 850, 835, 367, 831, 292, 651, 966, 268, 628, 925, 205, 824, 429, 917, 534, 323, 887, 3, 302, 134, 904, 300, 678, 929, 491, 229, 671, 817, 442, 678, 879, 373, 664, 990, 53, 395, 738, 570, 497, 113, 322, 557, 341, 641, 331, 932, 830, 433, 590, 738, 780, 50, 446, 504, 743, 311, 980, 88, 224, 732, 316, 664, 742, 69, 146, 801, 334, 41, 198, 629, 690, 869, 598, 612, 662, 385, 637, 769, 984, 316, 741, 980, 2, 794, 814, 730, 297, 503, 734, 836, 604, 674, 376, 692, 277, 727, 455, 975, 703, 115, 25, 552, 404, 460, 543, 738, 86, 488, 356, 929, 668, 835, 222, 413, 172, 221, 1000, 30, 888, 350, 514, 908, 870, 323, 991, 201, 738, 335, 189, 437, 604, 316, 514, 575, 531, 514, 318, 43, 592, 594, 9, 773, 609, 952, 708, 868, 291, 962, 572, 772, 291, 214, 992, 238, 275, 36, 882, 631, 376, 150, 838, 376, 862, 996, 258, 545, 331, 907, 958, 925, 503, 1, 745, 559, 147, 617, 487, 185, 623, 287, 658, 340, 84, 835, 563, 168, 845, 401, 395, 928, 277, 136, 890, 276, 45, 806, 121, 264, 416, 417, 596, 208, 106, 738, 352, 995, 746, 731, 72, 258, 112, 885, 445, 165, 74, 847, 633, 343, 721, 237, 20, 91, 575, 410, 765, 274, 233, 738, 893, 999, 283, 104, 414, 981, 448, 761, 47, 48, 725, 459, 265, 318, 564, 353, 260, 896, 874, 563, 492, 710, 336, 952, 80, 195, 326, 311, 716, 167, 561, 556, 234, 680, 631, 112, 573, 248, 422, 130, 219, 134, 75, 722, 188, 221, 238, 193, 689, 63, 787, 657, 956, 214, 895, 657, 169, 349, 575, 577, 869, 64, 325, 187, 471, 535, 572, 39, 872, 966, 22, 232, 427, 501, 855, 239, 487, 263, 335, 645, 461, 973, 447, 923, 922, 788, 286, 610, 55, 708, 827, 250, 355, 481, 379, 322, 926, 796, 815, 2, 952, 268, 257, 61, 795, 364, 999, 535, 494, 664, 619, 711, 228, 411, 587, 292, 345, 671, 640, 231, 384, 859, 88, 640, 838, 904, 27, 235, 605, 766, 887, 23, 438, 816, 764, 91, 12, 324, 709, 411, 659, 405, 927, 769, 505, 259, 383, 714, 333, 652, 648, 663, 604, 596, 231, 114, 320, 955, 689, 626, 495, 758, 96, 848, 43, 189, 848, 656, 114, 475, 349, 148, 995, 467, 94, 519, 141, 125, 598, 738, 822, 701, 194, 46, 936, 332, 370, 764, 944, 711, 889, 568, 508, 186, 981, 48, 400, 69, 182, 698, 25, 526, 808, 272, 963, 451, 335, 883, 718, 199, 185, 437, 81, 987, 4, 274, 482, 263, 509, 584, 767, 141, 53, 365, 14, 657, 712, 837, 161, 378, 525, 313, 685, 183, 869, 202, 382, 339, 351, 686, 15, 667, 636, 756, 553, 848, 57, 740, 862, 962, 838, 410, 722, 409, 589, 891, 370, 520, 790, 880, 276, 478, 26, 459, 671, 728, 301, 296, 75, 194, 173, 116, 938, 933, 977, 812, 863, 868, 286, 973, 984, 265, 631, 456, 436, 683, 28, 126, 319, 285, 62, 247, 88, 60, 824, 710, 26, 602, 897, 765, 998, 610, 138, 773, 555, 153, 114, 932, 21, 111, 171, 282, 246, 909, 419, 647, 781, 166, 966, 200, 521, 188, 808, 295, 685, 1000, 890, 353, 301, 983, 862, 527, 974, 241, 705, 437, 523, 213, 704, 421, 225, 428, 310, 255, 719, 243, 962, 757, 27, 476, 181, 138, 95, 309, 122, 500, 846, 627, 371, 470, 759, 255, 373, 520, 748, 856, 459, 71, 431, 782, 307, 524, 644, 130, 120, 56, 406, 387, 435, 201, 7, 392, 922, 503, 578, 331, 827, 954, 21, 351, 869, 65, 300, 697, 908, 505, 315, 198, 744, 892, 510, 307, 985, 129, 634, 773, 343, 640, 702, 748, 973, 594, 271, 151, 254, 513, 339, 843, 425, 153, 19, 309, 489, 333, 944, 442, 904, 447, 239, 487, 6, 230, 988, 656, 716, 488, 779, 362, 738, 663, 516, 432, 964, 142, 823, 353, 175, 797, 645, 613, 553, 26, 41, 946, 47, 479, 181, 964, 901, 251, 843, 715, 211, 366, 335, 16, 103, 547, 171, 276, 29, 165, 993, 424, 274, 334, 754, 982, 63, 963, 904, 150, 342, 301, 238, 152, 314, 892, 498, 958, 192, 806, 208, 681, 703, 970, 688, 5, 809, 705, 182, 230, 658, 531, 793, 303, 475, 825, 924, 538, 488, 100, 655, 524, 569, 655, 430, 808, 820, 402, 852, 760, 691, 751, 779, 868, 247, 688, 545, 780, 350, 400, 550, 307, 577, 803, 527, 302, 916, 984, 829, 257, 172, 392, 41, 233, 241, 587, 159, 176, 904, 926, 540, 324, 918, 177, 817, 585, 722, 89, 987, 476, 637, 210, 980, 905, 911, 547, 762, 490, 197, 718, 774, 982, 484, 781, 675, 152, 144, 412, 255, 800, 480, 901, 892, 309, 382, 873, 469, 662, 375, 499, 646, 436, 410, 866, 440, 708, 613, 842, 663, 604, 555, 133, 77, 458, 66, 660, 504, 635, 896, 621, 126, 995, 506, 7, 283, 11, 610, 11, 727, 667, 101, 589, 309, 240, 508, 368, 830, 805, 4, 259, 936, 39, 510, 645, 772, 993, 530, 932, 393, 19, 82, 915, 994, 853, 683, 183, 797, 61, 292, 942, 434, 846, 265, 316, 991, 751, 579, 182, 162, 454, 5, 194, 97, 451, 906, 177, 761, 988, 314, 425, 5, 63, 127, 565, 427, 774, 66, 195, 627, 731, 750, 586, 874, 599, 878, 759, 807},
			515062},
	}
	for _, tt := range tests {
		if got := maxProfit2(tt.arg); got != tt.want {
			t.Errorf("%v maxProfit() =  %v, want %v", tt.name, got, tt.want)
		}
	}
}
