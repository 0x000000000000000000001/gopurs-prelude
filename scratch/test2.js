const numSub = function (n1) { return function (n2) { return n1 - n2; }; };
const negate = function(a) { return numSub(0.0)(a); };
console.log(1.0 / negate(0.0));
