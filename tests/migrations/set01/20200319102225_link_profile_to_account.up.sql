ALTER TABLE `profiles`
  ADD CONSTRAINT profiles_account_id_accounts_fk
  FOREIGN KEY (account_id) REFERENCES `accounts`(id);
