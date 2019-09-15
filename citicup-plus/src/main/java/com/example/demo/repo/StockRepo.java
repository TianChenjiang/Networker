package com.example.demo.repo;

import com.example.demo.model.Stock;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;

import java.util.List;


public interface StockRepo extends JpaRepository<Stock,Long>,
        JpaSpecificationExecutor<Stock> {
    List<Stock> findAllByTsCode(String tscode);
}
