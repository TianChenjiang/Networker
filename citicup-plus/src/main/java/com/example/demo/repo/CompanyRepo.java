package com.example.demo.repo;

import com.example.demo.model.Company;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;

public interface CompanyRepo extends JpaRepository<Company, Long>,
        JpaSpecificationExecutor<Company> {
}
