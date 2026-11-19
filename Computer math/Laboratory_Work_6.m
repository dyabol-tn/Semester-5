%1
A = [1 2 3; 4 5 6; 7 8 9];
A(1,1) = 16;
disp(A);

%2
A = [1 2 3; 4 5 6; 7 8 9];
A(3,:) = 100;
disp(A);

%3
A = [1 2 3; 4 5 6; 7 8 9];
A(1,:) = [];
disp('Без первой строки:'); disp(A);

B = [1 2 3; 4 5 6; 7 8 9];
B(:,2) = [];
disp('Без 2-го столбца:'); disp(B);

%4
A = [1 2 3; 4 5 6; 7 8 9; 10 11 12];
A(1:2,:) = [];
disp(A);

%5
A = [1 2 3 4; 5 6 7 8; 9 10 11 12];
A(:, [2 4]) = [];
disp(A);

%6
A = [4 2 1; 3 5 2; 1 1 3];
B = [2 1 0; 1 3 1; 0 1 2];

detA = det(A);
invA = inv(A);
trA  = A';
trB  = B';

Sum  = A + B;
Prod = A * B;

DivR = A / B;
DivL = A \ B;

DivElem = A ./ B;

disp('det(A) = '); disp(detA);
disp('inv(A):');   disp(invA);
disp('A'':');      disp(trA);
disp('A+B:');      disp(Sum);
disp('A*B:');      disp(Prod);
disp('A/B:');      disp(DivR);
disp('A\B:');      disp(DivL);

%7
A = [1 2; 3 4];
A4 = A.^4;
disp(A4);

%8
x = -2*pi : 0.01 : 2*pi;
y = sin(x);

figure;
plot(x, y, 'g', 'LineWidth', 2);
xlabel('X');
ylabel('Y');
title('График функции y = sin(x)');
grid on;