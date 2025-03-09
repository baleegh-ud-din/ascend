package database

func CreateSchemas() error {
	// ENUM
	enums := []struct {
		name   string
		values []string
	}{
		// {"approval_status", []string{"Pending", "Approved", "Requesting Change", "Rejected"}},
	}

	for _, enum := range enums {
		if err := CreateEnum(enum.name, enum.values); err != nil {
			return err
		}
	}

	// TABLES
	tables := []struct {
		name   string
		schema string
	}{
		// {"approvals", `
		// 	urid BIGSERIAL PRIMARY KEY,
		// 	entity VARCHAR(255) NOT NULL,
		// 	item_id INTEGER NULL,
		// 	action approval_action NOT NULL,
		// 	status approval_status NOT NULL,
		// 	previous_data JSONB NULL,
		// 	updated_data JSONB,
		// 	message TEXT,
		// 	desflow VARCHAR(20),
		// 	updater VARCHAR(255) NOT NULL,
		// 	approver VARCHAR(255),
		// 	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		// 	last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		// `},
	}

	for _, table := range tables {
		if err := CreateTable(table.name, table.schema); err != nil {
			return err
		}
	}

	// INDEXES
	indexes := []struct {
		name      string
		tableName string
		column    string
	}{
		// {"idx_approvals_urid", "approvals", "urid"},
	}

	for _, index := range indexes {
		if err := CreateIndex(index.name, index.tableName, index.column); err != nil {
			return err
		}
	}

	return nil
}
