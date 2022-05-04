
CREATE TABLE "INVENTORY" (
	loc varchar(10) NOT NULL,
	lot varchar(10) NOT NULL,
	sku text NOT NULL,
	qty numeric(22, 5) NOT NULL,
	qtyallocated numeric(22, 5) NOT NULL,
	qtypicked numeric(22, 5) NOT NULL,
	qtyexpected numeric(22, 5) NOT NULL,
	create_date date NOT NULL DEFAULT now()
);


CREATE TABLE "ITM_MOVEMENT" (
	"type" text NOT NULL,
	to_loc varchar(10) NOT NULL,
	from_loc varchar(10) NOT NULL,
	to_lot varchar(10) NOT NULL,
	from_lot varchar(10) NOT NULL,
	sku text NOT NULL,
	qty numeric(22, 5) NOT NULL,
	create_date date NOT NULL DEFAULT now()
);



CREATE TABLE "LOC" (
	loc varchar(10) NOT NULL,
	width numeric NOT NULL DEFAULT 0,
	height numeric NOT NULL DEFAULT 0,
	length numeric NOT NULL DEFAULT 0,
	"cube" numeric NOT NULL DEFAULT 0,
	"type" text NOT NULL,
	flag text NOT NULL,
	status text NOT NULL,
	create_date date NOT NULL DEFAULT now(),
	CONSTRAINT "LOC_pkey" PRIMARY KEY (loc)
);


CREATE TABLE "LOT" (
	lot varchar(10) NOT NULL,
	lottable01 text NOT NULL,
	lottable02 text NOT NULL,
	lottable03 text NOT NULL,
	lottable04 text NOT NULL,
	expired_date date NOT NULL,
	create_date date NOT NULL DEFAULT now(),
	CONSTRAINT "LOT_pkey" PRIMARY KEY (lot)
);


CREATE TABLE "SKU" (
	sku text NOT NULL,
	whseid varchar(30) NOT NULL,
	def_uom varchar(50) NOT NULL,
	productcode text NOT NULL,
	serialkey text NOT NULL,
	"desc" text NOT NULL,
	"enable" bool NOT NULL DEFAULT true,
	create_date date NOT NULL DEFAULT now(),
	CONSTRAINT "SKU_pkey" PRIMARY KEY (sku),
	CONSTRAINT sku_unique UNIQUE (sku)
);


CREATE TABLE "UOM" (
	uom varchar(50) NOT NULL,
	"desc" text NOT NULL,
	create_date date NOT NULL DEFAULT now(),
	CONSTRAINT "UOM_pkey" PRIMARY KEY (uom),
	CONSTRAINT uom_unique UNIQUE (uom)
);

CREATE TABLE "UOM_CONVERSION" (
	id bigserial NOT NULL,
	base_uom varchar(50) NOT NULL,
	target_uom varchar(50) NOT NULL,
	base_qty numeric(22, 5) NOT NULL,
	target_qty numeric(22, 5) NOT NULL,
	create_date date NOT NULL DEFAULT now(),
	CONSTRAINT "UOM_CONVERSION_pkey" PRIMARY KEY (id),
	CONSTRAINT uom_conversion_unique UNIQUE (base_uom, target_uom)
);


CREATE TABLE "WHSE" (
	whseid varchar(30) NOT NULL,
	"desc" text NOT NULL,
	create_date date NOT NULL DEFAULT now(),
	CONSTRAINT "WHSE_pkey" PRIMARY KEY (whseid)
);

