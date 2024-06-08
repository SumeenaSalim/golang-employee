module Employees {
    type Employees {
        required name: str;
        position: str;
        salary: float64;
        required created_at: datetime {
            readonly := true;
            default := datetime_of_statement();
        };
        updated_at: datetime {
            rewrite update using (datetime_of_statement());
        };
    }
}