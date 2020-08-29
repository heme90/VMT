1.Purpose of Project

golang ddd 연습용 토이 프로젝트

2. Project domain<br/>
음료 자판기의 판매기록을 축적하여 고객사에서 해당 데이터를<br/>
사용, 상품기획, 판촉 등에 활용할수 있도록 구성

3. Project Output<br/>
외부와 소통 가능한  api 서버

4. Quality Assertion<br/>
PostMan API Tester 를 통한 테스트

5. UseCases, Actors<br/>
- 상품 판매 로깅(loggingGoods)<br/>
    상품 판매시 서버에 요청하여 해당 상품 판매정보 저장<br/>
    저장하는 정보는 다음과 같을 예정<br/>
    *상품명<br/>
    *가격<br/>
    *상품 카테고리<br/>
    *판매 시각<br/>
    *지역<br/>
    @상품 카테고리는 음료, 생필품,.. 등으로 추가되는 상수로 관리
    지역, 시각 등은 판촉에 필요한 정보<br/>
    지역을 남기지 말고 자판기를 엔티티로 등록하여 해당 자판기의 위치를 남기면?
    자판기의 판매정보에서는 구성상품 등도 중요함<br/>
    상품 카테고리도 자판기 카테고리로 변경가능<br/>
    
- 상품 판매 조회<br/>
    마케팅 등 판촉에 사용할 정보이므로 정해진 조회 수단보다는 조회 가능한 대쉬보드를 열어주는쪽이 좋을듯<br/>
    대쉬보드 연동 가능하도록 api 공개<br/>
    대쉬보드는 db에 꽂아도 되지않냐?<br/>
    쿼리툴을 공개하자<br/>
    body에 raw query 받아서 조회문일 경우 조회 가능하게 + 인증키도 받을수 있으면 좋음<br/>
    
    
6. 입출력 정의<br/>
    6-1.Persistence<br/>
    DB 통해서 관리할 예정<br/>
    DBMS 선택 --> VMT는 비영리 목적 개인 프로젝트이고 DBMS 에서 제공해주는 기능이 많았으면 좋겠음<br/>
    1. PostGreSQL<br/>
    2. NoSQL<br/>
    
    6-2. Protocol<br/> 
    액터가 사람인 경우와 머신인 경우, 등등 추가될 상황이 많으므로 rpc, http 대응 가능한 grpc 선택<br/>
     
    
    
    
    
    
    
    
    
    
    