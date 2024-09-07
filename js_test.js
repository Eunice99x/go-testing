let arr1 = [1, 2, 3, 4, 5];
let res = 0;

for (let i = 0; i < arr1.length; i++) {
	if (arr1[0] != arr1[i]) {
		res += arr1[i];
	}
}

console.log(res);
