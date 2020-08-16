1.Purpose of Project

golang ddd 연습용 토이 프로젝트

2. Project domain
음료 자판기의 판매기록을 축적하여 고객사에서 해당 데이터를
사용, 상품기획, 판촉 등에 활용할수 있도록 구성

3. Project Output
외부와 소통 가능한  api 서버

4. Quality Assertion
PostMan API Tester 를 통한 테스트

5. UseCases, Actors
- 상품 판매 로깅(loggingGoods)
    상품 판매시 서버에 요청하여 해당 상품 판매정보 저장
    저장하는 정보는 다음과 같을 예정
    1. 상품명
    2. 가격
    3. 상품 카테고리
    4. 판매 시각
    5. 지역
    *상품 카테고리는 음료, 생필품,.. 등으로 추가되는 상수로 관리
    지역, 시각 등은 판촉에 필요한 정보
    지역을 남기지 말고 자판기를 엔티티로 등록하여 해당 자판기의 위치를 남기면?
    자판기의 판매정보에서는 구성상품 등도 중요함
    상품 카테고리도 자판기 카테고리로 변경가능
    
- 상품 판매 조회
    마케팅 등 판촉에 사용할 정보이므로 정해진 조회 수단보다는 조회 가능한 대쉬보드를 열어주는쪽이 좋을듯
    대쉬보드 연동 가능하도록 api 공개
    대쉬보드는 db에 꽂아도 되지않냐?
    쿼리툴을 공개하자
    body에 raw query 받아서 조회문일 경우 조회 가능하게 + 인증키도 받을수 있으면 좋음
    
    
    
    
    
    