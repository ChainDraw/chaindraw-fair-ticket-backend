--
CREATE DATABASE IF NOT EXISTS chaindraw_fair_ticket;

USE chaindraw_fair_ticket;

-- 用户表
CREATE TABLE if NOT EXISTS tb_user (
    id BIGINT auto_increment primary key,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    identity_document_type int NOT NULL,
    identity_document_number VARCHAR(50) NOT NULL,
    identity_document_country VARCHAR(50) NOT NULL,
    identity_document_image_url VARCHAR(255) NOT NULL,
    wallet_address VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    create_at BIGINT NOT NULL,
    update_at BIGINT NOT NULL
)ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- 演唱会票表
CREATE table if NOT EXISTS tb_ticket (
    id BIGINT auto_increment primary key,
    concert_id VARCHAR(255) NOT NULL comment '演唱会id',
    ticket_type int NOT NULL comment '门票种类唯一键',
    type_name VARCHAR(255) NOT NULL comment '门票种类名称',
    num INT NULL comment '门票数量',
    price decimal(10, 2) NOT NULL comment '门票价格',
    ticket_img VARCHAR(255) collate utf8mb4_bin null comment '门票照片',
    trade INT NULL comment '是否可以二手交易',
    max_quantity_per_wallet int NOT NULL,
    create_at BIGINT NOT NULL,
    update_at BIGINT NOT NULL
)ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 演唱会表
CREATE table if NOT EXISTS tb_concert (
    id BIGINT auto_increment primary key,
    concert_name VARCHAR(255) NOT NULL,
    address VARCHAR(255) collate utf8mb4_bin NOT NULL comment '演唱会地址',
    concert_img_url VARCHAR(255) NOT NULL,
    lottery_start_date BIGINT NULL comment '订阅抽选开始时间',
    lottery_end_date BIGINT NULL comment '订阅抽选结束时间',
    concert_date BIGINT NOT NULL,
    concert_status int NOT NULL comment '演唱会状态 0未开始；1已过期',
    review_status int NOT NULL comment '审核状态 0 待审核； 1审核通过；',
    cancel_reason VARCHAR(255) comment '退票理由',
    create_at BIGINT NOT NULL,
    update_at BIGINT NOT NULL,
    remark VARCHAR(255) collate utf8mb4_bin null comment '备注描述'
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 监听事件表
CREATE TABLE `event_escrow_created` (
                                        `id` bigint NOT NULL AUTO_INCREMENT,
                                        `concert_id` varchar(255) NOT NULL,
                                        `ticket_type` varchar(255) NOT NULL,
                                        `escrow_address` varchar(255) NOT NULL,
                                        `created_at` datetime NOT NULL,
                                        `updated_at` datetime NOT NULL,
                                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `event_lotteryescrow__claimedfund` (
                                                    `id` bigint NOT NULL AUTO_INCREMENT,
                                                    `concert_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                    `ticket_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                    `organizer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                    `winner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                    `money` varchar(255) NOT NULL,
                                                    `created_at` datetime NOT NULL,
                                                    `updated_at` datetime NOT NULL,
                                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `event_lotteryescrow__completedraw` (
                                                     `id` bigint NOT NULL AUTO_INCREMENT,
                                                     `lottery_address` varchar(255) NOT NULL,
                                                     `created_at` datetime NOT NULL,
                                                     `updated_at` datetime NOT NULL,
                                                     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `event_lotteryescrow__deposited` (
                                                  `id` bigint NOT NULL AUTO_INCREMENT,
                                                  `concert_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                  `ticket_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                  `buyer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                  `money` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                  `created_at` datetime NOT NULL,
                                                  `updated_at` datetime NOT NULL,
                                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `event_lotteryescrow__nonwinner` (
                                                  `id` bigint NOT NULL AUTO_INCREMENT,
                                                  `concert_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                  `ticket_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                  `non_winner` varchar(255) NOT NULL,
                                                  `money` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                  `created_at` datetime NOT NULL,
                                                  `updated_at` datetime NOT NULL,
                                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `event_lotteryescrow__refunded` (
                                                 `id` bigint NOT NULL AUTO_INCREMENT,
                                                 `concert_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                 `ticket_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                 `buyer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                 `money` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                                 `created_at` datetime NOT NULL,
                                                 `updated_at` datetime NOT NULL,
                                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `event_lotteryescrow__winner` (
                                               `id` bigint NOT NULL AUTO_INCREMENT,
                                               `concert_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                               `ticket_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                               `winner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                               `created_at` datetime NOT NULL,
                                               `updated_at` datetime NOT NULL,
                                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `event_nft_delisted` (
                                      `id` bigint NOT NULL AUTO_INCREMENT,
                                      `seller` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                      `lottery_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                      `token_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                      `created_at` datetime NOT NULL,
                                      `updated_at` datetime NOT NULL,
                                      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `event_nft_listed` (
                                    `id` bigint NOT NULL AUTO_INCREMENT,
                                    `seller` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                    `lottery_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                    `token_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                    `price` varchar(255) NOT NULL,
                                    `created_at` datetime NOT NULL,
                                    `updated_at` datetime NOT NULL,
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `event_nft_sold` (
                                  `id` bigint NOT NULL AUTO_INCREMENT,
                                  `buyer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                  `lottery_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                  `token_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                  `price` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                  `created_at` datetime NOT NULL,
                                  `updated_at` datetime NOT NULL,
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DELIMITER //

CREATE TRIGGER auto_increment_ticket_type
    BEFORE INSERT ON tb_ticket
    FOR EACH ROW
BEGIN
    DECLARE max_value INT;

    SELECT MAX(ticket_type) INTO max_value
    FROM tb_ticket;

    IF max_value IS NULL THEN
        SET NEW.ticket_type = 1;
    ELSE
        SET NEW.ticket_type = max_value + 1;
END IF;
END;//

DELIMITER ;

INSERT INTO chaindraw_fair_ticket.event_escrow_created (id, concert_id, ticket_type, escrow_address, created_at, updated_at) VALUES (4, '10010', '1', '0x1234567890abcdef1234567890abcdef12345678', '2024-05-27 14:42:44', '2024-05-27 14:42:44');
INSERT INTO chaindraw_fair_ticket.event_nft_listed (id, seller, lottery_address, token_id, price, created_at, updated_at) VALUES (3, 'seller3', '0x7890abcdef1234567890abcdef1234567890abcd', '3', '1 ETH', '2024-05-27 14:42:44', '2024-05-27 14:42:44');
INSERT INTO chaindraw_fair_ticket.event_nft_listed (id, seller, lottery_address, token_id, price, created_at, updated_at) VALUES (2, 'seller2', '0xabcdef1234567890abcdef1234567890abcdef12', '2', '0.75 ETH', '2024-05-27 14:42:44', '2024-05-27 14:42:44');
INSERT INTO chaindraw_fair_ticket.event_nft_listed (id, seller, lottery_address, token_id, price, created_at, updated_at) VALUES (1, 'seller1', '0x1234567890abcdef1234567890abcdef12345678', '1', '0.5 ETH', '2024-05-27 14:42:44', '2024-05-27 14:42:44');
INSERT INTO chaindraw_fair_ticket.tb_concert (id, concert_id, concert_name, address, concert_img_url, lottery_start_date, lottery_end_date, concert_date, concert_status, review_status, CREATE_at, update_at, remark) VALUES (4, '10010', '宇多田光', '日本', 'www.test.com', 1716820964000, 1716820964000, 1716820964000, 0, 1, 1716820964000, 1716820964000, '备注');
INSERT INTO chaindraw_fair_ticket.tb_ticket (id, concert_id, ticket_type, type_name, num, price, ticket_img, trade, max_quantity_per_wallet, CREATE_at, update_at) VALUES (4, '10010', 1, 'VIP', 100, 10.00, 'www.test.com', 1, 10, 1716820964000, 1716820964000);
INSERT INTO chaindraw_fair_ticket.tb_ticket (id, concert_id, ticket_type, type_name, num, price, ticket_img, trade, max_quantity_per_wallet, CREATE_at, update_at) VALUES (5, '10010', 2, 'NORMAL', 1000, 1.00, 'www.test.com', 1, 10, 1716820964000, 1716820964000);
