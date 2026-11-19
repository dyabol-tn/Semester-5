% 1
x = 0:0.1:10;
y = sin(x);

figure;
plot(x, y, 'r--*', 'LineWidth', 1.5, 'MarkerSize', 6);
xlabel('x');
ylabel('y = sin(x)');
title('График функции y = sin(x)');
grid on;  

%2
x = 0:0.1:10;
figure;
plot(x, sin(x), 'r-',  'LineWidth', 1.5); hold on;
plot(x, cos(x), 'b--', 'LineWidth', 1.5);
plot(x, sin(x).*cos(x), 'g-.', 'LineWidth', 1.5);
hold off;
xlabel('x'); ylabel('y');
title('Несколько функций на одном графике');
legend('sin(x)', 'cos(x)', 'sin(x)cos(x)');
grid on;

%3
[X, Y] = meshgrid(-3:0.1:3, -3:0.1:3);
Z = sin(X).*cos(Y);

figure;
mesh(X, Y, Z);
shading interp;
colormap jet;
colorbar;
xlabel('x'); ylabel('y'); zlabel('z = sin(x)cos(y)');
title('Поверхность z = sin(x)cos(y)');

%4.1
v = [3, -1, 5, -7, 0, 2, -4, 8];
for k = 1:length(v)
    if v(k) < 0
        v(k) = 0;
    else
        v(k) = v(k)^2;
    end
end
disp(v);

%4.2
A = [1 2 3; 4 5 6; 7 8 9];
n = size(A, 1);
S = 0;
for i = 1:n
    for j = 1:n
        if j > i
            S = S + A(i, j);
        end
    end
end
fprintf('Сумма выше главной диагонали = %d\n', S);

%4.3
N = 7;
f = 1;
for k = 1:N
    f = f * k;
end
fprintf('%d! = %d\n', N, f);

%4.4
v = randi([-50, 50], 1, 10);
fprintf('Массив: '); disp(v);

sumPos = 0; cntNeg = 0; cntZero = 0;
minV = v(1); maxV = v(1);

for k = 1:length(v)
    if v(k) > 0
        sumPos = sumPos + v(k);
    elseif v(k) < 0
        cntNeg = cntNeg + 1;
    else
        cntZero = cntZero + 1;
    end
    if v(k) < minV, minV = v(k); end
    if v(k) > maxV, maxV = v(k); end
end

fprintf('Сумма положительных: %d\n', sumPos);
fprintf('Количество отрицательных: %d\n', cntNeg);
fprintf('Количество нулей: %d\n', cntZero);
fprintf('Минимум: %d, Максимум: %d\n', minV, maxV);

%4.5
v = [1, 10, -3, 25, 4, 8, 30, -7];
C = 5;
S = 0;
for k = 1:length(v)
    if v(k) > C
        S = S + v(k);
    end
end
fprintf('Сумма элементов > %d = %d\n', C, S);

%4.6
A = randi(9, 4, 4);
disp('Исходная матрица:'); disp(A);
[n, m] = size(A);
for i = 1:n
    for j = 1:m
        if j <= i
            A(i, j) = 0;
        end
    end
end
disp('После обнуления диагонали и ниже:'); disp(A);

%4.7
v = [4, -2, 7, 0, -5, 9, -1, 3];
sumPos = 0; cntNeg = 0;
for k = 1:length(v)
    if v(k) > 0
        sumPos = sumPos + v(k);
    elseif v(k) < 0
        cntNeg = cntNeg + 1;
    end
end
fprintf('Сумма положительных = %d\n', sumPos);
fprintf('Количество отрицательных = %d\n', cntNeg);