# Number Line Jumps

아래 과정으로 풀이가 진행 됨


문제를 공식화 하면 다음과 같음.

(v1 * i + x1) = (v2 * i + x2)

- i 는 점프 횟수
- x1, x2 는 캥거루의 시작 위치
- v1, v2 는 캥거루의 점프 거리

위 수식을 v는 v끼리, x는 x끼리 묶으면 아래와 같음
- (v1 - v2) = (x2 - x1) * i

각 항에 (x2 - x1) 를 나누면 아래와 같이 변함
- i = (v1 - v2) / (x2 - x1)

위 식에 대입해 볼 수 있는 것은 아래와 같음.
- i 는 0보다 큰 정수이다.

즉, 아래 2가지 식을 유추할 수 있다.
- (v1 - v2) / (x2 - x1) > 0
- (x2 - x1) % (v1 - v2) = 0

그렇기에, 위 2가지 조건을 만족하면 두 마리의 캥거루는 언젠가 같은 위치에 도달할 것임.

Problem : https://www.hackerrank.com/challenges/kangaroo/problem?isFullScreen=true