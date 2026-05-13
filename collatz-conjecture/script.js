function calculateTransformations(num) {
    var i = 0;
    var ifContinue = true;

    while (ifContinue) {
        if (num === 1) {
            ifContinue = false;
            break;
        } else if (num % 2 === 0) {
            num /= 2;
        } else {
            num *= 3;
            num += 1;
        }

        i++;
    }

    return i;
}

const n = 5;
result = calculateTransformations(n);
console.log(result);
