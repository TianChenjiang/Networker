package com.example.demo.repo;

import com.example.demo.model.Company;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import org.springframework.data.domain.Pageable;

import java.util.List;

public interface CompanyRepo extends JpaRepository<Company, Long>,
        JpaSpecificationExecutor<Company> {
    @Query(value = "select * from company c where c.ts_code in (select ts_code from stock s where s.name like %:key%)",nativeQuery = true)
    List<Company> getCompaniesByKey(
                                    @Param(value = "key") String key,Pageable pageable);
}
