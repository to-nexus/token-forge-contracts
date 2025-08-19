// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ITokenFactory {
    enum TokenType {
        ERC20,
        ERC721,
        ERC1155
    }

    // ========== 토큰 배포 함수 ==========

    /**
     * @dev ERC20 토큰 배포
     * @param owner 새 토큰의 소유자 EOA
     * @param manager 토큰 관리자 EOA (optional, 0이면 owner만 관리)
     * @param name 토큰 이름
     * @param symbol 토큰 심볼
     * @param decimals 토큰 소수점 자리수
     * @param initialSupply 초기 공급량 (optional, 0이면 민팅 안함)
     * @param initialRecipient 초기 공급량을 받을 주소 (optional, initialSupply가 0이 아닐 때 필수)
     * @param logic 사용할 로직 컨트랙트 주소 (optional, 0이면 기본 프리셋 사용)
     * @return tokenAddress 배포된 토큰 프록시 주소
     */
    function deployERC20(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory data,
        address logic
    ) external returns (address tokenAddress);

    /**
     * @dev ERC721 토큰 배포
     * @param owner 새 토큰의 소유자 EOA
     * @param manager 토큰 관리자 EOA (optional, 0이면 owner만 관리)
     * @param name 토큰 이름
     * @param symbol 토큰 심볼
     * @param baseTokenURI 기본 토큰 URI
     * @param logic 사용할 로직 컨트랙트 주소 (optional, 0이면 기본 프리셋 사용)
     * @return tokenAddress 배포된 토큰 프록시 주소
     */
    function deployERC721(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        string memory baseTokenURI,
        bytes memory data,
        address logic
    ) external returns (address tokenAddress);

    /**
     * @dev ERC1155 토큰 배포
     * @param owner 새 토큰의 소유자 EOA
     * @param manager 토큰 관리자 EOA (optional, 0이면 owner만 관리)
     * @param uri 토큰 URI 템플릿
     * @param logic 사용할 로직 컨트랙트 주소 (optional, 0이면 기본 프리셋 사용)
     * @return tokenAddress 배포된 토큰 프록시 주소
     */
    function deployERC1155(address owner, address manager, string memory uri, bytes memory data, address logic)
        external
        returns (address tokenAddress);
    //
    //    // ========== 토큰 등록 함수 (외부 배포된 토큰 등록) ==========
    //
    //    /**
    //     * @dev 외부에서 배포된 토큰을 서비스에 등록
    //     * @param service ForgeFactory에 등록된 서비스 이름
    //     * @param tokenType 토큰 타입 (20, 721, 1155)
    //     * @param tokenAddress 등록할 토큰 주소
    //     */
    //    function registerExternalToken(string calldata service, uint256 tokenType, address tokenAddress) external;

    // ========== 프리셋 로직 관리 함수 ==========

    /**
     * @dev 프리셋 로직 주소 추가/업데이트
     * @param tokenType 토큰 타입 (0:20, 1:721, 2:1155)
     * @param logicAddress 로직 컨트랙트 주소
     * @param add true면 추가, false면 제거
     */
    function setPresetLogics(TokenType tokenType, address[] calldata logicAddress, bool add) external;

    // ========== 조회 함수 ==========

    /**
     * @dev 프리셋 로직 주소 조회
     * @param tokenType 토큰 타입 (0:20, 1:721, 2:1155)
     * @return logicAddress 로직 컨트랙트 주소
     */
    function getPresetLogics(TokenType tokenType) external view returns (address[] memory);

    // ========== 관리자 함수 ==========

    //    /**
    //     * @dev ForgeFactory 주소 설정
    //     * @param forgeFactory ForgeFactory 컨트랙트 주소
    //     */
    //    function setForgeFactory(address forgeFactory) external;

    // ========== 이벤트 ==========

    event TokenDeployed(address indexed owner, TokenType indexed tokenType, address tokenAddress, address logicAddress);

    event PresetLogicSet(TokenType indexed tokenType, address indexed logicAddress);

    event PresetLogicRemoved(TokenType indexed tokenType, address indexed logicAddress);
}
