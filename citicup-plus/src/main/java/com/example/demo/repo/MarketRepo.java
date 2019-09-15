package com.example.demo.repo;

import com.example.demo.model.Market;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface MarketRepo extends JpaRepository<Market, Long>,
        JpaSpecificationExecutor<Market> {
    List<Market> findMarketsByTsCode(String tsCode);
}
