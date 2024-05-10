--
-- 用户表
CREATE TABLE tb_user
(
    id                          BIGINT       NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email                       VARCHAR(255) NOT NULL,
    phone                       VARCHAR(20)  NOT NULL,
    identity_document_type      INT NOT NULL, --0: 'passport', 1: 'id_card'
    identity_document_number    VARCHAR(50)  NOT NULL,
    identity_document_country   VARCHAR(50)  NOT NULL,
    identity_document_image_url VARCHAR(255) NOT NULL,
    wallet_address              VARCHAR(255) NOT NULL,
    password                    VARCHAR(255) NOT NULL,
    create_at                   BIGINT       NOT NULL,
    update_at                   BIGINT       NOT NULL
);


-- 演唱会表
CREATE TABLE tb_concert
(
    id              BIGINT       NOT NULL AUTO_INCREMENT PRIMARY KEY,
    concert_id      VARCHAR(255) NOT NULL,
    concert_name    VARCHAR(255) NOT NULL,
    concert_img_url VARCHAR(255) NOT NULL,
    concert_date    BIGINT       NOT NULL, -- 演唱会时间
    concert_status  INT      NOT NULL, --0: 未开始 1：已过期 2、已取消
    status          VARCHAR(255) NOT NULL,
    create_at       BIGINT       NOT NULL,
    update_at       BIGINT       NOT NULL
);

-- 演唱会票表
CREATE TABLE tb_ticket
(
    id                      BIGINT         NOT NULL AUTO_INCREMENT PRIMARY KEY,
    concert_id              VARCHAR(255)   NOT NULL,
    ticket_type             VARCHAR(255)   NOT NULL,
    type_name               INT            NOT NULL, -- 0: 普通 1: 高级 2: vip
    price                   DECIMAL(10, 2) NOT NULL,
    max_quantity_per_wallet INT            NOT NULL,
    create_at               BIGINT         NOT NULL,
    update_at               BIGINT         NOT NULL
);
