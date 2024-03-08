/*
 Navicat Premium Data Transfer

 Source Server         : postgres_localhost
 Source Server Type    : PostgreSQL
 Source Server Version : 160002 (160002)
 Source Host           : localhost:5432
 Source Catalog        : equipments
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 160002 (160002)
 File Encoding         : 65001

 Date: 09/03/2024 00:02:47
*/


-- ----------------------------
-- Table structure for equipments
-- ----------------------------
DROP TABLE IF EXISTS "public"."equipments";
CREATE TABLE "public"."equipments" (
  "id" int4 NOT NULL DEFAULT nextval('equipments_id_seq'::regclass),
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "model" varchar(255) COLLATE "pg_catalog"."default",
  "version" varchar(255) COLLATE "pg_catalog"."default",
  "expense" float8,
  "date_of_manufacture" int8,
  "color" varchar(255) COLLATE "pg_catalog"."default",
  "camera" varchar(255) COLLATE "pg_catalog"."default",
  "battery_life" varchar(255) COLLATE "pg_catalog"."default",
  "length" varchar(255) COLLATE "pg_catalog"."default",
  "width" varchar(255) COLLATE "pg_catalog"."default",
  "height" varchar(255) COLLATE "pg_catalog"."default",
  "weight" varchar(255) COLLATE "pg_catalog"."default",
  "material" varchar(255) COLLATE "pg_catalog"."default",
  "core" varchar(255) COLLATE "pg_catalog"."default",
  "sensor" varchar(255) COLLATE "pg_catalog"."default",
  "algorithm" varchar(255) COLLATE "pg_catalog"."default",
  "category" varchar(255) COLLATE "pg_catalog"."default",
  "language" varchar(255) COLLATE "pg_catalog"."default",
  "working_radius" varchar(255) COLLATE "pg_catalog"."default",
  "charging_voltage" varchar(255) COLLATE "pg_catalog"."default",
  "motor_type" varchar(255) COLLATE "pg_catalog"."default",
  "movement_maximumspeed" varchar(255) COLLATE "pg_catalog"."default",
  "control" varchar(255) COLLATE "pg_catalog"."default",
  "communication_interface" varchar(255) COLLATE "pg_catalog"."default",
  "bluetooth" varchar(255) COLLATE "pg_catalog"."default",
  "wireless" varchar(255) COLLATE "pg_catalog"."default",
  "manufacturer" varchar(255) COLLATE "pg_catalog"."default",
  "programming_language" varchar(255) COLLATE "pg_catalog"."default",
  "service_life" varchar(255) COLLATE "pg_catalog"."default",
  "applicable_equipment" varchar(255) COLLATE "pg_catalog"."default",
  "protocol" varchar(255) COLLATE "pg_catalog"."default",
  "type" varchar(255) COLLATE "pg_catalog"."default",
  "learning_speed" varchar(255) COLLATE "pg_catalog"."default",
  "success_rate" varchar(255) COLLATE "pg_catalog"."default",
  "hours_of_operation" varchar(255) COLLATE "pg_catalog"."default",
  "repetitive_positioning_accuracy" varchar(255) COLLATE "pg_catalog"."default",
  "screen" varchar(255) COLLATE "pg_catalog"."default",
  "other" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."equipments" OWNER TO "postgres";

-- ----------------------------
-- Primary Key structure for table equipments
-- ----------------------------
ALTER TABLE "public"."equipments" ADD CONSTRAINT "equipments_pkey" PRIMARY KEY ("id");
