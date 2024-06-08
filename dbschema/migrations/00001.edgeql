CREATE MIGRATION m1zkbsloifcjc5aqy23zgblsprfnrdzjargub37ezqnqvne4tr6nsa
    ONTO initial
{
  CREATE MODULE Employees IF NOT EXISTS;
  CREATE TYPE Employees::Employees {
      CREATE REQUIRED PROPERTY created_at: std::datetime {
          SET default := (std::datetime_of_statement());
          SET readonly := true;
      };
      CREATE REQUIRED PROPERTY name: std::str;
      CREATE PROPERTY position: std::str;
      CREATE PROPERTY salary: std::float64;
      CREATE PROPERTY updated_at: std::datetime {
          CREATE REWRITE
              UPDATE 
              USING (std::datetime_of_statement());
      };
  };
};
