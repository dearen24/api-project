package com.example.demo.service;
import com.example.demo.model.User;
import com.example.demo.repository.UserRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class UserService {
    private final UserRepository repository;

    public UserService(UserRepository repository) {
        this.repository = repository;
    }

    public List<User> getAllUsers() {
        return repository.findAll();
    }

    public User addUser(User user) {
        return repository.save(user);
    }

    public void deleteUser(Long id) {
        repository.deleteById(id);
    }

    public Optional<User> getUserById(Long id) {
        return repository.findById(id);
    }

    public Optional<User> findByUsername(String username) {
        return repository.findByUsername(username);
    }

    public User updateUser(Long id, User updatedUser) {
        return repository.findById(id)
                .map(user -> {
                    user.setUsername(updatedUser.getUsername());
                    user.setPassword_hash(updatedUser.getPassword_hash());
                    user.setEmail(updatedUser.getEmail());
                    user.setFirst_name(updatedUser.getFirst_name());
                    user.setLast_name(updatedUser.getLast_name());
                    user.setIs_active(updatedUser.getIs_active());
                    return repository.save(user);
                })
                .orElse(null);
    }
}
